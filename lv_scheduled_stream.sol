pragma solidity 0.4.24;

import {Content} from "./content.sol";
import {BaseContent} from "./base_content.sol";
//import {BaseLibrary} from "./base_library.sol";


/* -- Revision history --
LvRightsHolder20191007145500ML: First versioned released
*/
contract LvRightsHolder is Content {

    bytes32 public version = "LvRightsHolder20191007145500ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    struct Rights {
        bool active;
        bool canRecord;
        bool canPlayback;
    }


    mapping (address => Rights) public recordingStreams;
    mapping (bytes32 => Rights) public recordingPrograms; //Not used for now

    event ChangeStreamsRights(address stream, bool canRecord, bool canPlayback);
    event AuthorizeRecording(address stream, address accessor, bool decision);
    event AuthorizePlayback(address stream, address accessor, bool decision);

    //registering is schedule-based and the schedule has the same ownership as the stream, hence we can not ensure that the rights owner agrees
    function registerStream(address stream) public  {
        BaseContent streamObj = BaseContent(stream);
        require(streamObj.owner() == tx.origin);
        recordingStreams[stream] = Rights(true, true, true);
        emit ChangeStreamsRights(stream, true, true);
    }


    function canRecord(address stream, address accessor)  public view returns (bool) {
        return recordingStreams[stream].canRecord;
    }

    function authorizeRecording(address stream, address accessor)  public returns (bool) {
        bool decision = recordingStreams[stream].canRecord;
        LvScheduledRecording rec = LvScheduledRecording(msg.sender);
        if (rec.recordingAuth(address(this)) == 0) {
            emit AuthorizeRecording(stream, accessor, decision);
            rec.markRecordingAuth(decision);
            LvScheduledStream streamCustomContract =  LvScheduledStream(rec.recordingStreamContract());
            streamCustomContract.logRecordingAuthorization(decision);
        }
        return decision;
    }

    function canPlayback(address stream, address accessor) public view returns (bool) {
        return recordingStreams[stream].canPlayback;
    }

    function authorizePlayback(address stream, address accessor)  public returns (bool) {
        bool decision = recordingStreams[stream].canPlayback;
        LvScheduledRecording rec = LvScheduledRecording(msg.sender);
        if (rec.checkPlaybackAuth()) {
            emit AuthorizePlayback(stream, accessor, decision);
            rec.markPlaybackAuth(decision);
            LvScheduledStream streamCustomContract =  LvScheduledStream(rec.recordingStreamContract());
            streamCustomContract.logPlaybackAuthorization(decision);
        }
        return decision;
    }


    function changeStreamsRights(address stream, bool canRecord, bool canPlayback) public {
        Rights memory rights = recordingStreams[stream];
        require(rights.active);
        rights.canRecord = canRecord;
        rights.canPlayback = canPlayback;
        recordingStreams[stream] = rights;
        emit ChangeStreamsRights(stream, canRecord, canPlayback);
    }

    function enableRecording(address stream) public onlyOwner {
        Rights memory rights = recordingStreams[stream];
        require(rights.active);
        rights.canRecord = true;
        recordingStreams[stream] = rights;
        emit ChangeStreamsRights(stream, true, rights.canPlayback);
    }

    function disableRecording(address stream) public onlyOwner {
        Rights memory rights = recordingStreams[stream];
        require(rights.active);
        rights.canRecord = false;
        recordingStreams[stream] = rights;
        emit ChangeStreamsRights(stream, false, rights.canPlayback);
    }

}



/* -- Revision history --
LvSchStream20191007145700ML: First versioned released
*/

contract LvScheduledStream is Content {

    bytes32 public version = "LvSchStream20191007145700ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    uint public startTime;
    uint public endTime;
    bool public recordingEnabled; //stream level recording enable/disable (decision cascaded with rightsholder's)
    string public handle;

    address public recordingStream; //the content object for the recordable stream


    //mapping (address => address) public recordings; //maps recording objects to recordings Custom Contract

    address public schedule;

    event CreateRecording(address recObj, address recContract);
    event SetRecordingTimes(address recObj, uint recStartTime, uint recEndTime);
    event SetRecordingStatus(address recObj, string recStatus);
    event StreamwideRecordingBlock(address recObj);
    event EnableRecording();
    event DisableRecording();
    event recordingAuthorized(address rightsHolder, bool decision);
    event playbackAuthorized(address rightsHolder, bool decision);

    event StartStream();
    event StopStream();

    constructor() public payable {
        if (msg.sender != tx.origin) {
            recordingStream = msg.sender;
        }
        startTime = 0;
        endTime = 0;
        handle = "";
        recordingEnabled = true;
    }

    function setRecordingStream(address stream) public onlyOwner  { //only required if contract is instanciated manually, otherwise it is set correctly in the constructor
       recordingStream = stream;
    }

    // When a recordable stream is created a contract is created to track copies of that recording
    function runCreate() public payable returns (uint) {

        if (recordingStream == 0x0) {
            setRecordingStream(msg.sender);
            return 0;
        }

        return 10; //contract should be used only for one stream before being used to generate recordings
    }

    function enableRecording() public onlyOwner  {
        recordingEnabled = true;
        emit EnableRecording();
    }

    function disableRecording() public onlyOwner  {
        recordingEnabled = false;
        emit DisableRecording();
    }

    function createRecording(address baseContent, uint256 _startTime, uint256 _endTime) public returns (address) {
        address instanceAddress = new LvScheduledRecording();
        LvScheduledRecording rec = LvScheduledRecording(instanceAddress);
        rec.setTimes(_startTime, _endTime);
        if (baseContent != 0x0) {
            rec.setContentAddress(baseContent);
        }
        emit CreateRecording(baseContent, instanceAddress);
    }

    function setSchedule(address _schedule) public onlyOwner  {
        schedule = _schedule;
    }

    function startStream(string _handle) public onlyOwner  {
        handle = _handle;
        startTime = now;
        endTime = 0; //to allow re-opening
        emit StartStream();
    }

    function stopStream() public onlyOwner  {
        handle = "";
        endTime = now;
        emit StopStream();
    }

    function logRecordingStatus() public returns (uint8){
        LvScheduledRecording rec = LvScheduledRecording(msg.sender);
        uint8 recStatus = rec.recordingStatus();
        if ( recStatus == 10) {
            emit SetRecordingStatus(rec.contentAddress(), "recording");
            return recStatus;
        }
        if ( recStatus == 100) {
            emit SetRecordingStatus(rec.contentAddress(), "complete");
        }
        return recStatus;
    }

    function logRecordingTimes() public {
        LvScheduledRecording rec = LvScheduledRecording(msg.sender);
        emit SetRecordingTimes(rec.contentAddress(), rec.startTime(), rec.endTime());
    }

    function logRecordingAuthorization(bool decision) public {
        emit recordingAuthorized(msg.sender, decision);
    }

    function logPlaybackAuthorization(bool decision) public {
        emit playbackAuthorized(msg.sender, decision);
    }

    function logStreamwideBlock() public {
        LvScheduledRecording rec = LvScheduledRecording(msg.sender);
        emit StreamwideRecordingBlock(rec.contentAddress());
    }

}


/* -- Revision history --
LvSchRecording20191008192700ML: First versioned released
LvSchRecording20191017104300ML: Consider cases of mixed programs, some with rights holder some without
*/

contract LvScheduledRecording is Content {

    bytes32 public version ="LvSchRecording20191017104300ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    uint public startTime;
    uint public endTime;

    uint8 public recordingStatus; //0: not started, 10: started, 100: completed

    address public recordingStreamContract; //custom contract address for the stream object
    address public contentAddress; //contract address of recording base object

    address[] public rightsHolders;
    uint256 public rightHoldersLength;

    uint256 public rightsRegisteredUpto;


    mapping (bytes32 => uint8) public playbackAuth; //key is (accessRequestId, rightsHolder address) - 1 is authorized, 2 is not
    mapping (address => uint8) public recordingAuth; // key is rightsHolder address - 1 is authorized, 2 is not

    event SetTimes(uint startTime, uint endTime);
    event UpdateRecordingStatus(uint8 status);
    event AuthorizeRecording(bool decision);
    event AuthorizePlayback(bool decision);

    constructor() public payable {
        startTime = 0;
        endTime = 0;
        recordingStreamContract = msg.sender;
        recordingStatus = 0;
        rightHoldersLength = 0;
        rightsRegisteredUpto = 0;
    }

    function setContentAddress(address _contentAddress) public onlyOwner {
         contentAddress = _contentAddress;
         BaseContent contentObj = BaseContent(_contentAddress);
         contentObj.setContentContractAddress(address(this));
    }


   //keccak256(abi.encodePacked(request_ID, msg.sender))

    function addRightsHolder(address rightsHolder, uint256 _startTime, uint256 duration, bytes signature) public onlyOwner returns (bool){

        // ------------
        // TO BE ADDED: Check that the signature is valid and match the address of the stream owner
        // ------------
        bool signatureMatchStreamOwner = true;
        require(signatureMatchStreamOwner);


        if ((_startTime >= startTime) && ((endTime == 0) || (_startTime <= endTime))) {
            if ((rightsRegisteredUpto == 0) || (rightsRegisteredUpto == _startTime)) { //no gaps
                rightsRegisteredUpto = _startTime + duration;
                if (rightsHolder != 0x0) {
                    for (uint256 i=0; i < rightHoldersLength; i++) {
                        if (rightsHolders[i] == rightsHolder) {
                            return false; //only add rightsHolder if not present yet to avoid duplicates
                        }
                    }
                    rightsHolders.push(rightsHolder);
                    rightHoldersLength += 1;
                }
            }
        }
        return true;
    }


    //Same logic as authorizeRecording but as a view
    function canRecord(uint256 _startTime, uint256 duration) public view returns (bool) {
        if ((rightsRegisteredUpto != 0) && (rightsRegisteredUpto < (_startTime + duration))) {
            return false; //attempting to register beyond declared segments
        }
        LvScheduledStream stream = LvScheduledStream(recordingStreamContract);
        if (!stream.recordingEnabled()) {
            return false;
        }
        for (uint256 i=0; i < rightHoldersLength; i++) {
            LvRightsHolder rightsHolder = LvRightsHolder(rightsHolders[i]);
            if (rightsHolder.canRecord(stream.recordingStream(), tx.origin) == false) {
                return false;
            }
        }
        return true;
    }


    function authorizeRecording(uint256 _startTime, uint256 duration) public returns (bool) {
        if ((rightsRegisteredUpto != 0) && (rightsRegisteredUpto < (_startTime + duration))) {
            emit AuthorizeRecording(false);
            return false; //attempting to register beyond declared segments
        }
        LvScheduledStream stream = LvScheduledStream(recordingStreamContract);
        if (!stream.recordingEnabled()) {
            stream.logStreamwideBlock();
            emit AuthorizeRecording(false);
            return false;
        }
        for (uint256 i=0; i < rightHoldersLength; i++) {
            LvRightsHolder rightsHolder = LvRightsHolder(rightsHolders[i]);

            if (rightsHolder.authorizeRecording(stream.recordingStream(), tx.origin) == false) {
                emit AuthorizeRecording(false);
                return false;
            }
        }
        emit AuthorizeRecording(true);
        return true;
    }


    function canPlayback() public view returns (bool) {
        for (uint256 i=0; i < rightHoldersLength; i++) {
            LvRightsHolder rightsHolder = LvRightsHolder(rightsHolders[i]);
            LvScheduledStream stream = LvScheduledStream(recordingStreamContract);
            if (rightsHolder.canPlayback(stream.recordingStream(), tx.origin) == false) {
                return false;
            }
        }
        return true;
    }

    function authorizePlayback() public returns (bool) {
        for (uint256 i=0; i < rightHoldersLength; i++) {
            LvRightsHolder rightsHolder = LvRightsHolder(rightsHolders[i]);
            LvScheduledStream stream = LvScheduledStream(recordingStreamContract);
            if (rightsHolder.authorizePlayback(stream.recordingStream(), tx.origin) == false) {
                emit AuthorizePlayback(false);
                return false;
            }
        }
        emit AuthorizePlayback(true);
        return true;
    }

    function setStartTime(uint _startTime) public onlyOwner { //whether start and end time can be modified is debatable
        //modification is restricted to owner of recording content object or someone with editor privilege
        if (_startTime == 0) {
            // We could assume 0 means the beginning of the stream, but since we are assuming a single stream object, it does not make sense
            startTime = now;
        } else {
            startTime = _startTime;
        }
        emit SetTimes(startTime, endTime);
        LvScheduledStream stream = LvScheduledStream(recordingStreamContract);
        stream.logRecordingTimes();
    }

    function setEndTime(uint _endTime) public { //whether start and end time can be modified is debatable
        //modification is restricted to owner of recording content object or someone with editor privilege
        if (_endTime == 0) { //0 ends the recording
            endTime = now;
        } else {
            endTime = _endTime;
        }
        emit SetTimes(startTime, endTime);
        LvScheduledStream stream = LvScheduledStream(recordingStreamContract);
        stream.logRecordingTimes();
    }

    function setTimes(uint _startTime, uint _endTime) public {
        if (_startTime == 0) {
            startTime = now;
        } else {
            startTime = _startTime;
        }
        endTime = _endTime;
        emit SetTimes(startTime, endTime);
        LvScheduledStream stream = LvScheduledStream(recordingStreamContract);
        stream.logRecordingTimes();
    }

    function updateRecordingStatus(uint8 _recordingStatus) public onlyOwner {
        recordingStatus = _recordingStatus;
        emit UpdateRecordingStatus(recordingStatus);
        LvScheduledStream stream = LvScheduledStream(recordingStreamContract);
        stream.logRecordingStatus();
    }

    function markRecordingAuth(bool decision) public onlyOwner {
        if (decision) {
            recordingAuth[msg.sender] = 1;
        } else {
            recordingAuth[msg.sender] = 2;
        }
    }

    function checkPlaybackAuth() public view onlyOwner returns (bool) {
        uint256 requestID = BaseContent(contentAddress).requestID();
        return playbackAuth[keccak256(abi.encodePacked(requestID, msg.sender))] != 0;
    }

    function markPlaybackAuth(bool decision) public onlyOwner {
        uint256 requestID = BaseContent(contentAddress).requestID();

        if (decision) {
            playbackAuth[keccak256(abi.encodePacked(requestID, msg.sender))] = 1;
        } else {
            playbackAuth[keccak256(abi.encodePacked(requestID, msg.sender))] = 2;
        }
    }


    function runKill() public payable returns (uint) {
        return 100; //when base content object is destroyed, custom contract should be too
    }



}
