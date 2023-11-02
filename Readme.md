# steps to run the code

## Pre-requisite: need to have POSTGRESS running in your local

- install all dependencies

- try running test cases present inside clients

  - TestInsertTicker - used to do insert all intraday ticker
  - TestInsertAdjustedTicker - used to insert all adjusted ticker

- once both testcase becomes successful. try validating that

- update postgres config in start.sh

- start your server using `sh start.sh`

- you can find openapi's collection in docs. either use postman or you can use gcurl to open the api and their docs

## Note: api has reflection integrated.
