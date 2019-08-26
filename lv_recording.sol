pragma solidity 0.4.24;

import {Content} from "./content.sol";
import {BaseContent} from "./base_content.sol";
import {BaseLibrary} from "./base_library.sol";



/* -- Revision history --
LvRecStream20190812201700ML: First versioned released
LvRecStream20190823104800ML: Adding fields to store stream handle, and exposing start to stop time.
LvRecStream20190825165500ML: Adding stream-wide event logging of recordings.
*/


contract LvRecordableStream is Content {

    bytes32 public version = "LvRecStream20190825165500ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    uint public startTime;
    uint public endTime;
    bool public recordingEnabled;
    string public handle;

    address public recordingStream; //the content object for the recordable stream

    event CreateRecording(address recObj, address recContract);
    event SetRecordingTimes(address recObj, uint recStartTime, uint recEndTime);
    event SetRecordingStatus(address recObj, string recStatus);

    event StartStream();
    event StopStream();

    constructor() public payable {
        recordingStream = msg.sender;
        startTime = 0;
        endTime = 0;
        handle = "";
    }

    // When a recordable stream is created a contract is created to track copies of that recording
    function runCreate() public payable returns (uint) {

        if (recordingEnabled) {
            address instanceAddress = new LvRecording();
            LvRecording rec = LvRecording(instanceAddress);
            rec.setContentAddress(msg.sender);
            BaseContent obj = BaseContent(msg.sender);
            obj.setContentContractAddress(instanceAddress);
            emit CreateRecording(msg.sender, instanceAddress);
            return 0;
        }
        //emit LogBool("Recording is not enabled", recordingEnabled);
        return 0; // To allow for creation of Provider --- debatable. We might want to prevent use and instead create a factory contract used to create a new provider
    }

    function enableRecording() public onlyOwner  {
        recordingEnabled = true;
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
        LvRecording rec = LvRecording(msg.sender);
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
        LvRecording rec = LvRecording(msg.sender);
        emit SetRecordingTimes(rec.contentAddress(), rec.startTime(), rec.endTime());
    }

}


/* -- Revision history --
LvRecording20190812210100ML: First versioned released
LvRecording20190825165500ML: Adding stream-wide event logging of recordings.
*/





contract LvRecording is Content {

    bytes32 public version ="LvRecording20190825165500ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    uint public startTime;
    uint public endTime;

    uint8 public recordingStatus; //0: not started, 10: started, 100: completed

    address public recordingStreamContract;
    address public contentAddress;

    event SetTimes(uint startTime, uint endTime);
    event UpdateRecordingStatus(uint8 status);

    constructor() public payable {
        startTime = 0;
        endTime = 0;
        recordingStreamContract = msg.sender;
        recordingStatus = 0;
    }

    function setContentAddress(address _contentAddress) public onlyOwner {
         contentAddress = _contentAddress;
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
        LvRecordableStream stream = LvRecordableStream(recordingStreamContract);
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
        LvRecordableStream stream = LvRecordableStream(recordingStreamContract);
        stream.logRecordingTimes();
    }

    function setTimes(uint _startTime, uint _endTime) public {
        startTime = _startTime;
        endTime = _endTime;
        emit SetTimes(startTime, endTime);
        LvRecordableStream stream = LvRecordableStream(recordingStreamContract);
        stream.logRecordingTimes();
    }

    function updateRecordingStatus(uint8 _recordingStatus) public onlyOwner {
        recordingStatus = _recordingStatus;
        emit UpdateRecordingStatus(recordingStatus);
        LvRecordableStream stream = LvRecordableStream(recordingStreamContract);
        stream.logRecordingStatus();
    }


}
