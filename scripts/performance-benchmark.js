import http from 'k6/http';
import { check, sleep } from 'k6';

let accountCreated = false;
let accountId = null;
let baseUrl = 'http://localhost:3001';
let ledgerId = 1;

export let options = {
    stages: [
        { duration: '1m', target: 100 }, // Ramp up to 100 users over 1 minute
        { duration: '5m', target: 100 }, // Stay at 100 users for 5 minutes
        { duration: '1m', target: 0 } // Ramp down to 0 users over 1 minute
    ]
};

export default function () {

    // Perform transaction using POST on /transaction
    let transactionPayload = {
        amount: 100,
        debit_account_id: 1,
        credit_account_id: 2,
        ledger_id: ledgerId
    };
    let performTransaction = http.post(`${baseUrl}/transaction`, JSON.stringify(transactionPayload));
    check(performTransaction, {
        'Perform Transaction Successful': (res) => res.status === 200
    });

    // Add a sleep delay between API calls
    sleep(1);
}