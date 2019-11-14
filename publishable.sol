pragma solidity 0.4.24;

/**
 * Ownable
 * The Ownable contract stores owner and creator addresses, and provides basic authorization control
 * functions, this simplifies the implementation of "user permissions".
 */

/* -- Revision history --
Publishable20190221100500ML: First versioned released
*/


contract Publishable {

    bytes32 public version ="Publishable20190221100500ML"; //class name (max 16), date YYYYMMDD, time HHMMSS and Developer initials XX

    event SetStatusCode(int statusCode);

    int public statusCode; // 0: accessible, - in draft, + in review
    // application have discretion to make up their own status codes to represent their workflow

    function updateStatus(int status_code) public returns (int) {
        statusCode = status_code;
        emit SetStatusCode(statusCode);
        return statusCode;
    }


}
