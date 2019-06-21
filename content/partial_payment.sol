pragma solidity 0.4.24;

import {Content} from "../content.sol";
import {BaseContent} from "../base_content.sol";

//
// This contract implements an approval process in which after a 60%
// completion of the submission, a partial payment of 40% of the licensing
// fee is paid to the contributor, with the remaining balance paid upon
// full, final approval.
//

contract PartialPayment is Content {

    function runStatusChange(int proposed_status_code) public payable returns (int, int256)
    {
        BaseContent contentObj = BaseContent(msg.sender);
        uint8 percentComplete = contentObj.percentComplete();
        uint256 licensingFee = contentObj.licensingFee();
        uint256 licensingFeeReceived = contentObj.licensingFeeReceived();

        //if percentComplete is over 60%, status will be changed to in-review
        if (percentComplete >= 60) {
            if (percentComplete == 100) {
                if (proposed_status_code > 0) {
                    emit RunStatusChange(proposed_status_code, 1, 0);
                    return (1, 0); //set to review for final approval
                }

                if (proposed_status_code == 0) {
                  //calculate full payoff amount
                    int256 fullPayoff = int256(licensingFee - licensingFeeReceived);
                    emit RunStatusChange(proposed_status_code, 0, fullPayoff);
                    return (0, fullPayoff); //set to approved
                }
                //return to draft so no additional payoff
                emit RunStatusChange(proposed_status_code, proposed_status_code, 0);
                return (proposed_status_code, 0);
            }

            if (proposed_status_code > 0) {
                emit RunStatusChange(proposed_status_code, 3, 0);
                return (3, 0); //set to review for provisional approval
            }

            if (proposed_status_code == 0) {
              //calculate partial payoff amount
                uint256 partialPayoff = (licensingFee * 40) / 100;
                if (licensingFeeReceived > partialPayoff) {
                    emit RunStatusChange(proposed_status_code, -3, 0);
                    return (-3, 0); //set to approved draft
                } else {
                    emit RunStatusChange(proposed_status_code, -3, int256(partialPayoff - licensingFeeReceived));
                    return (-3, int256(partialPayoff - licensingFeeReceived)); //set to approved draft
                }
            }
            //return to draft so no additional payoff
            emit RunStatusChange(proposed_status_code, proposed_status_code, 0);
            return (proposed_status_code, 0);
        }
        emit RunStatusChange(proposed_status_code, -1, 0);
        return (-1, 0); //if partial has not reached 60%, then keep in draft and no payoff at all
    }
}
