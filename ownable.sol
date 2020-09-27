pragma solidity 0.4.24;

/**
 * Ownable
 * The Ownable contract stores owner and creator addresses, and provides basic authorization control
 * functions, this simplifies the implementation of "user permissions".
 */

/* -- Revision history --
Ownable20190221100500ML: First versioned released
Ownable20190315141500ML: Migrated to 0.4.24
Ownable20190528193800ML: Added contentSpace as all descendant objects need it anyway
Ownable20200210110100ML: Added versionAPI, set to 3.0
*/

interface IAdmin {
    function isAdmin(address _adminAddr) external view returns (bool);
}

contract Ownable {

    bytes32 public version ="Ownable20200210110100ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX
    bytes32 public versionAPI = "3.0";
    address public creator;
    address public owner;
    address public contentSpace;

    // TODO: don't know if we can easily change this to msg.sender...
    //  since object in the content hierarchy are created through factories, msg.sender would end up being the factory
    //  instead of the end-user as expected. it theory we could change the space to set the owner as msg.sender but
    //  we currently can't change spaces.
    // Moreover, this likely is safe because I'm not sure what purpose a spoofing attack would have that ends up with
    //  the spoofee as the owner of the object ...?
    constructor() public payable {
        creator = tx.origin;
        owner = msg.sender;
    }

    // "Fallback" function - necessary if this contract needs to be paid
    function () public payable { }

    /**
     * Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    modifier onlyCreator() {
        require(msg.sender == creator);
        _;
    }

    /**
     * Allows the current owner to transfer control of the contract to a newOwner.
     *  newOwner: The address to transfer ownership to.
     */
    function transferOwnership(address newOwner) public onlyOwner {
        require(newOwner != address(0));
        owner = newOwner;
    }

    function transferCreatorship(address newCreator) public onlyCreator {
        require(newCreator != address(0));
        creator = newCreator;
    }

    function kill() public onlyOwner {
        selfdestruct(owner);  // kills contract; send remaining funds back to owner
    }
}

contract Adminable is Ownable, IAdmin {
    // meant to be overridden in derived classes
    function isAdmin(address _candidate) public view returns (bool) {
        if (_candidate == owner) {
            return true;
        }
        return false;
    }

    modifier onlyAdmin() {
        require(isAdmin(msg.sender));
        _;
    }
}
