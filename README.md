# vending-system

This is a simulation of centralized system which controls several vending machines. Each vending machine can serve coffee or soft drinks, and before notifying by the customer, it stays in the idle state. After the customer selects, the machine switches to working mode, and the selected product is delivered after payment.

To use this machines, we need to import the "machine" package in our main function and create one or more VendingMachine objects using the NewVendingMachine constructor function. The NewVendingMachine function takes three arguments: the machine ID, the initial quantity of coffee, and the initial quantity of soft drinks. Main function, also, will be our controlling system.

## usage

1. Select a vending machine number from 1 to 3 when prompted.
2. Choose the type of beverage you want to purchase (coffee or soft drink).
3. Insert a coin by typing "1" and pressing enter.
4. Wait for your beverage to be dispensed.


## local-run

To run this project locally, 

cd src
go run .


## docker-run

Create image and run it,

cd src
docker build --tag vending-system:latest .
docker run -it vending-system
