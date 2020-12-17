package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type FarmtoFork struct {
}

type Farmer struct {
	ObjectType     string    `json:"docType"`
	FarmerName     string    `json:"farmerName"`
	FarmerId       uint64    `json:"farmerId"`
	WalletBalance  uint64    `json:"walletBalance"`
}

type Produce struct {
	FarmerId       uint64 `json:"farmerId"`
	ProduceName    string `json:"produceName"`
	Weight         uint64 `json:"weight"`
	ProduceId      uint64 `json:"produceId"`
	AggreId        uint64 `json:"aggreId"`
	ProduceAmt     uint64 `json:"produceAmt"`
	Paidstatus     string `json:"paidstatus"`
	Deliverystatus string `json:"deliverystatus"`
}

type Aggregator struct {
	ObjectType         string `json:"docType"`
	AggrName string `json:"aggrName"` 
	AggrId uint64 `json:"aggrId"` 
	WalletBalance uint64 `json:"walletBalance"`
}


type Transporter struct{
	TransportName string `json:"transportName"` 
	WalletBalance uint64 `json:"walletBalance"`
}

type Inventory struct {
	AggrId string `json:"aggrId"`
	ProduceId string `json:"produceId"`
	InvCount uint64 `json:"invCount"`
}

type BidProduce struct{
	Produceid uint64 `json:"Produceid"` 
	Weight uint64 `json:"weight"`  
	BidId uint64 `json:"bidId"`
	AggrId uint64 `json:"aggreId"`
	WholeSalerId uint64 `json:"wholeSalerId"`
	BidAmt uint64 `json:"bidAmt"`
	Bidstatus string `json:"bidstatus"`
	Paidstatus string `json:"paidstatus"`
	Deliverystatus string `json:"deliverystatus"`
	Transportpaystatus string `json:"transportpaystatus"`
}

type Wholesaler struct{
	ObjectType         string `json:"docType"`
	WholesalerName string `json:"wholesalerName"` 
	WholeSalerId uint64 `json:"wholeSalerId"` 
	WalletBalance uint64 `json:"walletBalance"`
}

// ============================================================
// creatwholesalekey - create unique key for aggr id
// ============================================================
func creatwholesalekey(_id string) string {
	var targetkey string
	targetkey = "W" + _id 
	return targetkey
}

func createbidkey(_id string) string {
	var targetkey string
	targetkey = "B" + _id
	return targetkey
}


// ============================================================
// createaggrkey - create unique key for aggr id
// ============================================================
func createaggrkey(_aggrid string) string {
	var targetkey string
	targetkey = "A" + _aggrid 
	return targetkey
}

// ============================================================
// createinvkey - create unique key to count inventory for an aggr id
// ============================================================
func createinvkey(_aggrid,_produceid string) string {
	var targetkey string
	targetkey = "I" + _aggrid +_produceid
	return targetkey
}

// ============================================================
// createfarmerkey - create unique key for farmer id
// ============================================================
func createfarmerkey(_id string) string {
	var targetkey string
	targetkey = "F" + _id
	return targetkey
}

// ============================================================
// createproducekey - create unique key for produce id
// ============================================================
func createproducekey(_id string) string {
	var targetkey string
	targetkey = "P" + _id
	return targetkey
}

// ============================================================
// createtransportkey - create unique key for transport
// ============================================================
func createtransportkey() string {
	var targetkey string
	targetkey = "T"
	return targetkey
}


func (s *FarmtoFork) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// type FabricSetup struct{
// 	ChannelID         string `json:"ChannelID"`
// 	ChainCodeID string `json:"ChainCodeID"` 
// }

func main() {
	// Create a new Smart Contract
	err := shim.Start(new(FarmtoFork))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}

	// fSetup := FabricSetup{
    //     // Network parameters
    //    // OrdererID: "orderer.hf.chainhero.io",

    //     // Channel parameters
    //     ChannelID:     "mychannel",
    //     //ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/chainHero/heroes-service/fixtures/artifacts/chainhero.channel.tx",

    //     // Chaincode parameters
    //     ChainCodeID:     "FarmtoFork",
    //     // ChaincodeGoPath: os.Getenv("GOPATH"),
    //     // ChaincodePath:   "github.com/chainHero/heroes-service/chaincode/",
    //     // OrgAdmin:        "Admin",
    //     // OrgName:         "org1",
    //     // ConfigFile:      "config.yaml",

    //     // User parameters
    //     //UserName: "User1",
	// }
	
	//Create a api
	router := mux.NewRouter()

    //GetPerson and CreatePerson are the functions implemented in the chaincode.

    router.HandleFunc("/api/getFarmer",getFarmerweb).Methods("GET")
   // router.HandleFunc("/api/registerFarmer", registerFarmer).Methods("POST")

    log.Fatal(http.ListenAndServe(":8000", router))
}

func getFarmerweb(w http.ResponseWriter, r *http.Request){
	log.Println("all farmers are here")
}

func (s *FarmtoFork) Invoke(stub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "registerFarmer" {
		return s.registerFarmer(stub, args)
	} else if function == "getFarmer" {
		return s.getFarmer(stub, args)
	} else if function == "registerProduce" {
		return s.registerProduce(stub, args)
	} else if function == "getProduce" {
		return s.getProduce(stub, args)
	} else if function == "registerAggregator" {
		return s.registerAggregator(stub, args)
	} else if function == "getAggregator"{
		return s.getAggregator(stub, args)
	} else if function == "markDeliveryFarmer"{
		return s.markDeliveryFarmer(stub, args)
	} else if function == "getInventory"{
		return s.getInventory(stub, args)
	} else if function == "registerWholesaler" {
		return s.registerWholesaler(stub, args)
	} else if function == "getWholesaler"{
		return s.getWholesaler(stub, args)
	} else if function == "registerBid"{
		return s.registerBid(stub, args)
	} else if function == "getBid"{
		return s.getBid(stub, args)
	} else if function == "approveBid"{
		return s.approveBid(stub, args)
	} else if function == "registerTransporter"{
		return s.registerTransporter(stub, args)
	} else if function == "markDeliveryTransport"{
		return s.markDeliveryTransport(stub, args)
	}  else if function == "getTransporter"{
		return s.getTransporter(stub, args)
	}  else if function == "markDeliveryAggr"{
		return s.markDeliveryAggr(stub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

// ============================================================================================================================
//  registerFarmer - Create function to register farmer details in the blockchain (input-> farmer name, aadhar id)
// ============================================================================================================================

func (s *FarmtoFork) registerFarmer(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var err error
	var farmerkey string

	fmt.Println("running registerFarmer()")

	if len(args) != 2 {
		fmt.Println("Incorrect number of arguments. Expecting 'farmer id' and 'farmer name'")
		return shim.Error("Incorrect number of arguments. Expecting 'farmer id' and 'farmer name'")
	}

	farmerkey = createfarmerkey(args[0])

	f := Farmer{}
	f.ObjectType = "Farmer"
	
	f.FarmerId, err = strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		return shim.Error(err.Error())
	}
	f.FarmerName = args[1]

	farmerAsBytes, err := stub.GetState(farmerkey)
	if err != nil {
		return shim.Error(err.Error())
	} else if farmerAsBytes != nil {
		fmt.Println("Farmer already exists: " + f.FarmerName)
		return shim.Error("Farmer already exists: " + f.FarmerName)
	}

	farmerDetAsBytes, err := json.Marshal(f)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(farmerkey, farmerDetAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// ============================================================================================================================
//  registerProduce - Create function to register farmer produce details in the blockchain
// ============================================================================================================================

func (s *FarmtoFork) registerProduce(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var err error
	var farmerkey string
	var producekey string

	fmt.Println("running registerProduce()")

	if len(args) != 8 {
		fmt.Println("Incorrect number of arguments. Expecting 8")
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}

	farmerkey = createfarmerkey(args[0])
	farmerAsBytes, err := stub.GetState(farmerkey)
	if err != nil {
		return shim.Error(err.Error())
	} else if farmerAsBytes == nil {
		fmt.Println("Farmer id: " + args[0] + " does not exist")
		return shim.Error("Farmer id: " + args[0] + " does not exist")
	}

	producekey = createproducekey(args[3])
	pAsBytes, err := stub.GetState(producekey)
	if err != nil {
		return shim.Error(err.Error())
	} else if pAsBytes != nil {
		fmt.Println("Produce: " + args[1] + " exists")
		return shim.Error("Produce: " + args[1] + "exists")
	}

	p := Produce{}
	
	p.FarmerId,err = strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		fmt.Println("Argument should be a number")
		return shim.Error(err.Error())
	}
	p.ProduceName = args[1]
	p.Weight,err = strconv.ParseUint(args[2], 10, 64)
	if err != nil {
		fmt.Println("Argument should be a number")
		return shim.Error(err.Error())
	}
	p.ProduceId,err = strconv.ParseUint(args[3], 10, 64)
	if err != nil {
		fmt.Println("Argument should be a number")
		return shim.Error(err.Error())
	}
	p.AggreId,err = strconv.ParseUint(args[4], 10, 64)
	if err != nil {
		fmt.Println("Argument should be a number")
		return shim.Error(err.Error())
	}
	p.ProduceAmt,err = strconv.ParseUint(args[5], 10, 64)
	if err != nil {
		fmt.Println("Argument should be a number")
		return shim.Error(err.Error())
	}
	p.Paidstatus = "PENDING"
	p.Deliverystatus = "PENDING"

	produceasBytes, _ := json.Marshal(p)
	err = stub.PutState(producekey, produceasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// ============================================================================================================================
//  getFarmer - Query function to return farner details in the blockchain (input-> farmer id)
// ============================================================================================================================
func (s *FarmtoFork) getFarmer(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var farmerkey string
	var farmerid string

	fmt.Println("running getFarmer()")

	if len(args) != 1 {
		fmt.Println("Incorrect number of arguments. Expecting Farmer id")
		return shim.Error("Incorrect number of arguments. Expecting Farmer id")
	}

	farmerid = args[0]
	farmerkey = createfarmerkey(farmerid)
	farmerAsBytes, err := stub.GetState(farmerkey)
	if err != nil {
		return shim.Error(err.Error())
	} else if farmerAsBytes == nil {
		fmt.Println("Farmer id: " + args[0] + " does not exist")
		return shim.Error("Farmer id: " + args[0] + " does not exist")
	}

	// f := Farmer{}
	// json.Unmarshal(farmerAsBytes, &f)

	// farmerDetAsBytes, _ := json.Marshal(f)
	// fmt.Println("Farmer data: " + string(farmerDetAsBytes))
	// return shim.Success(farmerDetAsBytes)
	fmt.Println("Farmer data: " + string(farmerAsBytes))
	return shim.Success(farmerAsBytes)
}

// ============================================================================================================================
//  getProduce - Query function to return produce details in the blockchain 
// ============================================================================================================================
func (s *FarmtoFork) getProduce(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var producekey string

	fmt.Println("running getProduce()")

	if len(args) != 1 {
		fmt.Println("Incorrect number of arguments. Expecting Produce id")
		return shim.Error("Incorrect number of arguments. Expecting Produce id")
	}

	producekey = createproducekey(args[0])
	produceAsBytes, err := stub.GetState(producekey)
	if err != nil {
		return shim.Error(err.Error())
	} else if produceAsBytes == nil {
		fmt.Println("Produce id: " + args[0] + " does not exist")
		return shim.Error("Produce id: " + args[0] + " does not exist")
	}

	p := Produce{}
	json.Unmarshal(produceAsBytes, &p)

	produceDetAsBytes, _ := json.Marshal(p)
	fmt.Println("Produce data: " + string(produceDetAsBytes))
	return shim.Success(produceDetAsBytes)
}

// ============================================================================================================================
//  registerAggregator - Create function to register aggregator details in the blockchain (input-> aggr name, aadhar id)
// ============================================================================================================================

func (s *FarmtoFork) registerAggregator(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var err error
	var aggrkey string 
	
	fmt.Println("running registerAggregator()")
	
		if len(args) != 2 {
			fmt.Println("Incorrect number of arguments.Expecting 'Aggregator id' and 'Aggregator name'")
			return shim.Error("Incorrect number of arguments. Expecting 'Aggregator id' and 'Aggregator name'")
		}
	
		aggrkey = createaggrkey(args[0])
		aggrAsBytes, err := stub.GetState(aggrkey)
		if err != nil {
			return shim.Error(err.Error())
		} else if aggrAsBytes != nil {
			fmt.Println("Aggregator  id: " + args[0] + " exists")
			return shim.Error("Aggregator id: " + args[0] + " exists")
		}
	
		f := Aggregator{}
		f.ObjectType = "Aggregator"
		f.AggrId,err =strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			return shim.Error(err.Error())
		}
		f.AggrName = args[1]
		f.WalletBalance = 1000
	
		aggrDetAsBytes, _ := json.Marshal(f)
		
		err = stub.PutState(aggrkey, aggrDetAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
	
		return shim.Success(nil)
}
	
// ============================================================================================================================
//  markDeliveryFarmer - Update function to mark delivery of farmer's produce (input-> aggr id, farmer id, produce id, payment amount)
// ============================================================================================================================
	
func (s *FarmtoFork) markDeliveryFarmer(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var err error
	var aggrid string 
	var farmerid string 
	var produceid string 
	var aggrkey string 
	//var pay_amount uint64
	
	fmt.Println("running markDeliveryFarmer()")
		
	if len(args) != 3 {
			return shim.Error("Incorrect number of arguments. Expecting 3")
	}
		
	aggrid = args[0]
	farmerid = args[1]
	produceid = args[2]
	//pay_amount,_ = strconv.ParseUint(args[3], 10, 64)
	aggrkey = createaggrkey(aggrid)
	aggrAsBytes, err := stub.GetState(aggrkey)
	if err != nil {
			return shim.Error(err.Error())
	} else if aggrAsBytes == nil {
			fmt.Println("Aggregator id: " + aggrid + " does not exist")
			return shim.Error("Aggregator id: " + aggrid + " does not exist")
	}

	a := Aggregator{}
	json.Unmarshal(aggrAsBytes, &a)

	farmerkey := createfarmerkey(farmerid)
	farmerAsBytes, err := stub.GetState(farmerkey)
	if err != nil {
			return shim.Error(err.Error())
	} else if farmerAsBytes == nil {
			fmt.Println("Farmer id: " + farmerid + " does not exist")
			return shim.Error("Farmer id: " + farmerid+ " does not exist")
	}

	f := Farmer{}
	json.Unmarshal(farmerAsBytes, &f)
	
	invkey := createinvkey(aggrid,produceid)
	inv := Inventory{}
	
	
	producekey := createproducekey(produceid)
	produceAsBytes, err := stub.GetState(producekey)
	if err != nil {
		return shim.Error(err.Error())
	} else if produceAsBytes == nil {
		fmt.Println("Produce id: " + produceid + " does not exist")
		return shim.Error("Produce id: " + produceid + " does not exist")
	}

	p := Produce{}
	json.Unmarshal(produceAsBytes, &p)
	if p.Deliverystatus != "DELIVERED"{
					p.Deliverystatus = "DELIVERED"
	}

	if p.Paidstatus != "PAID"{
	a.WalletBalance = a.WalletBalance - p.ProduceAmt

			
		if a.WalletBalance < 0{
			return shim.Error("Aggregator id: " + aggrid + " does not have enough balance for payment to farmer: "+farmerid)
		} else{
			p.Paidstatus = "PAID"
			f.WalletBalance = f.WalletBalance + p.ProduceAmt
			inv.AggrId = aggrid
			inv.ProduceId = produceid
			inv.InvCount = inv.InvCount + p.Weight

			aggrDetAsBytes, _ := json.Marshal(a)
			err = stub.PutState(aggrkey, aggrDetAsBytes)
			if err != nil {
				return shim.Error(err.Error())
			}	

			farmerDetAsBytes, _ := json.Marshal(f)
			err = stub.PutState( "F" + farmerid, farmerDetAsBytes)
			if err != nil {
					return shim.Error(err.Error())
			}
				
			invDetAsBytes, _ := json.Marshal(inv)
			err = stub.PutState(invkey, invDetAsBytes)
			if err != nil {
					return shim.Error(err.Error())
			}
		}
	}	
	
	produceDetAsBytes, _ := json.Marshal(p)
	err = stub.PutState(producekey, produceDetAsBytes)
	if err != nil {
			return shim.Error(err.Error())
	}
	
	return shim.Success(nil)
}
	
	// ============================================================================================================================
	//  getInventory - Query function to return inventory details for aggregator in the blockchain (input-> aggr id)
	// ============================================================================================================================
	func (s *FarmtoFork) getInventory(stub shim.ChaincodeStubInterface, args []string) sc.Response {
		var invkey string
		var aggrid string
		var produceid string
		fmt.Println("running getInventory()")
	
		if len(args) != 2 {
			fmt.Println("Incorrect number of arguments. Expecting Aggregator id and produce id")
			return shim.Error("Incorrect number of arguments. Expecting Aggregator id and produce id")
		}
	
		//Get hdr
		aggrid = args[0]
		produceid = args[1]
		invkey = createinvkey(aggrid,produceid)
		invAsBytes, err := stub.GetState(invkey)
		if err != nil {
			return shim.Error(err.Error())
		} else if invAsBytes == nil {
			fmt.Println("Aggregator id: " + aggrid + " does not have any Inventory for produce: "+produceid)
			return shim.Error("Aggregator id: " + aggrid + " does not have any Inventory for produce: "+produceid)
		}
	
		f := Inventory{}
		json.Unmarshal(invAsBytes, &f)
	
		invDetAsBytes, _ := json.Marshal(f)
		fmt.Println("Inventory data: " + string(invDetAsBytes))
		return shim.Success(invDetAsBytes)
	}
	
	// ============================================================================================================================
	//  getAggregator - Query function to return aggreator details in the blockchain (input-> aggr id)
	// ============================================================================================================================
	func (s *FarmtoFork) getAggregator(stub shim.ChaincodeStubInterface, args []string) sc.Response {
		var aggrkey string
		var aggrid string
	
		fmt.Println("running getAggregator()")
	
		if len(args) != 1 {
			fmt.Println("Incorrect number of arguments. Expecting aggr id")
			return shim.Error("Incorrect number of arguments. Expecting aggr id")
		}
	
		//Get hdr
		aggrid = args[0]
		aggrkey = createaggrkey(aggrid)
		aggrAsBytes, err := stub.GetState(aggrkey)
		if err != nil {
			return shim.Error(err.Error())
		} else if aggrAsBytes == nil {
			fmt.Println("Aggregator  id: " + aggrid + " does not exist")
			return shim.Error("Aggregator id: " + aggrid + " does not exist")
		}
	
		f := Aggregator{}
		json.Unmarshal(aggrAsBytes, &f)
	
		aggrDetAsBytes, _ := json.Marshal(f)
		fmt.Println("Aggregator data: " + string(aggrDetAsBytes))
		return shim.Success(aggrDetAsBytes)
	}

	// ============================================================================================================================
//  registerWholesaler - Create function to register Wholesaler details in the blockchain 
// ============================================================================================================================

func (s *FarmtoFork) registerWholesaler(stub shim.ChaincodeStubInterface, args []string) sc.Response {
var err error
var wholekey string 

fmt.Println("running registerWholesaler()")

	if len(args) != 2 {
		fmt.Println("Incorrect number of arguments.Expecting 'Wholesaler id' and 'wholesaler name'")
		return shim.Error("Incorrect number of arguments. Expecting 'Wholesaler id' and 'wholesaler name'")
	}

	wholekey = creatwholesalekey(args[0])
	wholeAsBytes, err := stub.GetState(wholekey)
	if err != nil {
		return shim.Error(err.Error())
	} else if wholeAsBytes != nil {
		fmt.Println("Wholesaler  id: " + args[0] + " exists")
		return shim.Error("Wholesaler id: " + args[0] + " exists")
	}

	f := Wholesaler{}
	f.ObjectType = "Wholesaler"
	f.WholeSalerId,err =strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		return shim.Error(err.Error())
	}
	f.WholesalerName = args[1]
	f.WalletBalance = 1000

	wholeDetAsBytes, _ := json.Marshal(f)
	
	err = stub.PutState(wholekey, wholeDetAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// ============================================================================================================================
//  registerBid - Create fucntoin for wholesaler to bid the produces
// ============================================================================================================================

func (s *FarmtoFork) registerBid(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var err error
	var wholeid string 
	var aggrid string 
	var produceid string 
	var bidid string 
	var wholekey string 
	var aggrkey string 
	var producekey string 
	var bidkey string 
	var bid_amount uint64
	var weight uint64

	fmt.Println("running registerBid()")
	
	if len(args) != 10{
		return shim.Error("Incorrect number of arguments. Expecting 10")
	}
	
	wholeid = args[0]
	aggrid = args[1]
	produceid = args[2]
	bidid = args[3]
	weight,err = strconv.ParseUint(args[4], 10, 64)
	if err != nil {
		return shim.Error(err.Error())
	}
	bid_amount,err = strconv.ParseUint(args[5], 10, 64)
	if err != nil {
		return shim.Error(err.Error())
	}
	wholekey = creatwholesalekey(wholeid)
	wholeAsBytes, err := stub.GetState(wholekey)
	if err != nil {
		return shim.Error(err.Error())
	} else if wholeAsBytes == nil {
		fmt.Println("Wholesaler id: " + wholeid + " does not exist")
		return shim.Error("Wholesaler id: " + wholeid + " does not exist")
	}

	aggrkey = createaggrkey(aggrid)
	aggrAsBytes, err := stub.GetState(aggrkey)
	if err != nil {
			return shim.Error(err.Error())
	} else if aggrAsBytes == nil {
			fmt.Println("Aggregator id: " + aggrid + " does not exist")
			return shim.Error("Aggregator id: " + aggrid + " does not exist")
	}

	producekey = createproducekey(produceid)
	produceAsBytes, err := stub.GetState(producekey)
	if err != nil {
		return shim.Error(err.Error())
	} else if produceAsBytes == nil {
		fmt.Println("Produce id: " + produceid + " does not exist")
		return shim.Error("Produce id: " + produceid+ " does not exist")
	}

	bidkey = createbidkey(bidid)
	bAsBytes, err := stub.GetState(bidkey)
	if err != nil {
		return shim.Error(err.Error())
	} else if bAsBytes != nil {
		fmt.Println("Bid: " + bidid + " exists")
		return shim.Error("Bid: " + bidid + "exists")
	}

	b := BidProduce{}
	
	b.WholeSalerId,err = strconv.ParseUint(wholeid, 10, 64)
	if err != nil {
		return shim.Error(err.Error())
	} 
	b.AggrId,err = strconv.ParseUint(aggrid, 10, 64)
	if err != nil {
		return shim.Error(err.Error())
	} 
	b.Produceid,err = strconv.ParseUint(produceid, 10, 64)
	if err != nil {
		return shim.Error(err.Error())
	} 
	b.Weight = weight
	b.BidId,err = strconv.ParseUint(bidid, 10, 64)
	if err != nil {
		return shim.Error(err.Error())
	} 
	b.BidAmt = bid_amount
	b.Bidstatus = "PENDING"
	b.Paidstatus = "PENDING"
	b.Deliverystatus = "PENDING"
	b.Transportpaystatus = "PENDING"

	bidDetAsBytes, _ := json.Marshal(b)
	err = stub.PutState(bidkey, bidDetAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// ============================================================================================================================
//  approveBid - Update fucntoin for aggregator to approve bids from wholesaler
// ============================================================================================================================

func (s *FarmtoFork) approveBid(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var err error
	var wholeid string 
	var bidid string 
	var wholekey string 
	var bidkey string 

	fmt.Println("running approveBid()")
	
	if len(args) != 2{
		fmt.Println("Incorrect number of arguments. Expecting Wholesaler id and bid id")
		return shim.Error("Incorrect number of arguments. Expecting Wholesaler id and bid id")
	}
	
	wholeid = args[0]
	bidid = args[1]
	
	wholekey = creatwholesalekey(wholeid)
	wholeAsBytes, err := stub.GetState(wholekey)
	if err != nil {
		return shim.Error(err.Error())
	} else if wholeAsBytes == nil {
		fmt.Println("Wholesaler id: " + wholeid + " does not exist")
		return shim.Error("Wholesaler id: " + wholeid + " does not exist")
	}

	bidkey = createbidkey(bidid)
	bAsBytes, err := stub.GetState(bidkey)
	if err != nil {
		return shim.Error(err.Error())
	} else if bAsBytes == nil {
		fmt.Println("Bid: " + bidid + " does not exist")
		return shim.Error("Bid: " + bidid + " does not exist")
	}

	b := BidProduce{}
	json.Unmarshal(bAsBytes, &b)
	
	if b.Bidstatus == "PENDING" {
		b.Bidstatus = "ACCEPTED"
		bidDetAsBytes, _ := json.Marshal(b)
		err = stub.PutState(bidkey, bidDetAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
	}else{
		return shim.Error("Bid: " + bidid + " is not pending for approval. Hence, Transaction Rejected!")
	}

	return shim.Success(nil)
}

	// ============================================================================================================================
//  registerTransporter - Create function to register Transporter details in the blockchain 
// ============================================================================================================================

func (s *FarmtoFork) registerTransporter(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var err error
	var tkey string 
	
	fmt.Println("running registerTransporter()")
	
	tkey = createtransportkey()
		tAsBytes, err := stub.GetState(tkey)
		if err != nil {
			return shim.Error(err.Error())
		} else if tAsBytes != nil {
			fmt.Println("Transporter exists")
			return shim.Error("Transporter exists")
		}
	
		t := Transporter{}
		t.TransportName = "Kumble"
		tdetAsBytes, _ := json.Marshal(t)
		err = stub.PutState(tkey, tdetAsBytes)
		if err != nil {
		return shim.Error(err.Error())
		}

	return shim.Success(nil)
}

// ============================================================================================================================
//  markDeliveryTransport - Update fucntoin for Transporter to mark delivery of goods
// ============================================================================================================================

func (s *FarmtoFork) markDeliveryTransport(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var err error
	var wholeid string 
	var aggrid string 
	var bidid string 
	var wholekey string 
	var bidkey string 
	var aggrkey string 
	var tkey string 
	var pay_amount uint64

	fmt.Println("running markDeliveryTransport()")
	
	if len(args) != 4{
		fmt.Println("Incorrect number of arguments. Expecting 4")
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	
	wholeid = args[0]
	aggrid = args[1]
	bidid = args[2]
	pay_amount,err= strconv.ParseUint(args[3], 10, 64)
	if err != nil {
		return shim.Error(err.Error())
	}

	wholekey = creatwholesalekey(wholeid)
	wholeAsBytes, err := stub.GetState(wholekey)
	if err != nil {
		return shim.Error(err.Error())
	} else if wholeAsBytes == nil {
		fmt.Println("Wholesaler id: " + wholeid + " does not exist")
		return shim.Error("Wholesaler id: " + wholeid + " does not exist")
	}

	aggrkey = createaggrkey(aggrid)
	aggrAsBytes, err := stub.GetState(aggrkey)
	if err != nil {
			return shim.Error(err.Error())
	} else if aggrAsBytes == nil {
			fmt.Println("Aggregator id: " + aggrid + " does not exist")
			return shim.Error("Aggregator id: " + aggrid + " does not exist")
	}

	a := Aggregator{}
	json.Unmarshal(aggrAsBytes, &a)

	bidkey = createbidkey(bidid)
	bAsBytes, err := stub.GetState(bidkey)
	if err != nil {
		return shim.Error(err.Error())
	} else if bAsBytes == nil {
		fmt.Println("Bid: " + bidid + " does not exist")
		return shim.Error("Bid: " + bidid + " does not exist")
	}

	b := BidProduce{}
	json.Unmarshal(bAsBytes, &b)
	
	tkey = createtransportkey()
	tAsBytes, err := stub.GetState(tkey)
	if err != nil {
		return shim.Error(err.Error())
	} else if tAsBytes == nil {
		fmt.Println("Transporter does not exist")
		return shim.Error("Transporter does not exist")
	}
	t := Transporter{}
	json.Unmarshal(tAsBytes, &t)

	if b.Deliverystatus == "PENDING" {
		b.Deliverystatus = "DELIVERED"
		if b.Transportpaystatus == "PENDING"{
			b.Transportpaystatus = "PAID"
			t.WalletBalance = t.WalletBalance + pay_amount
			a.WalletBalance = a.WalletBalance - pay_amount

			tDetAsBytes, _ := json.Marshal(t)
			err = stub.PutState(tkey, tDetAsBytes)
			if err != nil {
				return shim.Error(err.Error())
			}

			aggrDetAsBytes, _ := json.Marshal(a)
			err = stub.PutState(aggrkey, aggrDetAsBytes)
			if err != nil {
				return shim.Error(err.Error())
			}
		}

		bidDetAsBytes, _ := json.Marshal(b)
		err = stub.PutState(bidkey, bidDetAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
	}else{
		return shim.Error("Bid: " + bidid + " is not pending for delivery. Hence, Transaction Rejected!")
	}

	return shim.Success(nil)
}

// ============================================================================================================================
//  getWholesaler - Query function to return wholesaler details in the blockchain
// ============================================================================================================================
func (s *FarmtoFork) getWholesaler(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var wholekey string
	var wholeid string

	fmt.Println("running getWholesaler()")

	if len(args) != 1 {
		fmt.Println("Incorrect number of arguments. Expecting wholesaler id")
		return shim.Error("Incorrect number of arguments. Expecting wholesaler id")
	}

	//Get hdr
	wholeid = args[0]
	wholekey = creatwholesalekey(wholeid)
	wholeAsBytes, err := stub.GetState(wholekey)
	if err != nil {
		return shim.Error(err.Error())
	} else if wholeAsBytes == nil {
		fmt.Println("Wholesaler id: " + wholeid + " does not exist")
		return shim.Error("Wholesaler id: " + wholeid + " does not exist")
	}

	f := Wholesaler{}
	json.Unmarshal(wholeAsBytes, &f)

	wholeDetAsBytes, _ := json.Marshal(f)
	fmt.Println("Wholesaler data: " + string(wholeDetAsBytes))
	return shim.Success(wholeDetAsBytes)
}

// ============================================================================================================================
//  getTransporter - Query function to return wholesaler details in the blockchain
// ============================================================================================================================
func (s *FarmtoFork) getTransporter(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var tkey string

	fmt.Println("running getTransporter()")

	//Get hdr
	tkey = createtransportkey()
	tAsBytes, err := stub.GetState(tkey)
	if err != nil {
		return shim.Error(err.Error())
	} else if tAsBytes == nil {
		fmt.Println("Transporter does not exist")
		return shim.Error("Transporter does not exist")
	}

	f := Transporter{}
	json.Unmarshal(tAsBytes, &f)

	tDetAsBytes, _ := json.Marshal(f)
	fmt.Println("Transpoter data: " + string(tDetAsBytes))
	return shim.Success(tDetAsBytes)
}

// ============================================================================================================================
//  getBid - Query function to return wholesaler bid details from blockchain
// ============================================================================================================================
func (s *FarmtoFork) getBid(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var wholekey string
	var bidkey string
	var wholeid string
	var bidid string

	fmt.Println("running getBid()")

	if len(args) != 2 {
		fmt.Println("Incorrect number of arguments. Expecting wholesaler id and bid id")
		return shim.Error("Incorrect number of arguments. Expecting wholesaler id and bid id")
	}

	//Get hdr
	wholeid = args[0]
	bidid = args[1]
	wholekey = creatwholesalekey(wholeid)
	wholeAsBytes, err := stub.GetState(wholekey)
	if err != nil {
		return shim.Error(err.Error())
	} else if wholeAsBytes == nil {
		fmt.Println("Wholesaler id: " + wholeid + " does not exist")
		return shim.Error("Wholesaler id: " + wholeid + " does not exist")
	}

	bidkey = createbidkey(bidid)
	bAsBytes, err := stub.GetState(bidkey)
	if err != nil {
		return shim.Error(err.Error())
	} else if bAsBytes == nil {
		fmt.Println("Bid: " + bidid + " does not exist")
		return shim.Error("Bid: " + bidid + " does not exist")
	}

	b := BidProduce{}
	json.Unmarshal(bAsBytes, &b)

	bDetAsBytes, _ := json.Marshal(b)
	fmt.Println("Bid data: " + string(bDetAsBytes))
	return shim.Success(bDetAsBytes)
}

// ============================================================================================================================
//  markDeliveryAggr - Update fucntoin for Wholesaler to mark delivery of goods and pay aggregator
// ============================================================================================================================

func (s *FarmtoFork) markDeliveryAggr(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	var err error
	var wholeid string 
	var aggrid string 
	var bidid string 
	var wholekey string 
	var bidkey string 
	var aggrkey string 

	fmt.Println("running markDeliveryAggr()")
	
	if len(args) != 3{
		fmt.Println("Incorrect number of arguments. Expecting 3")
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	
	wholeid = args[0]
	aggrid = args[1]
	bidid = args[2]
		
	wholekey = creatwholesalekey(wholeid)
	wholeAsBytes, err := stub.GetState(wholekey)
	if err != nil {
		return shim.Error(err.Error())
	} else if wholeAsBytes == nil {
		fmt.Println("Wholesaler id: " + wholeid + " does not exist")
		return shim.Error("Wholesaler id: " + wholeid + " does not exist")
	}
	w := Wholesaler{}
	json.Unmarshal(wholeAsBytes, &w)

	aggrkey = createaggrkey(aggrid)
	aggrAsBytes, err := stub.GetState(aggrkey)
	if err != nil {
			return shim.Error(err.Error())
	} else if aggrAsBytes == nil {
			fmt.Println("Aggregator id: " + aggrid + " does not exist")
			return shim.Error("Aggregator id: " + aggrid + " does not exist")
	}

	a := Aggregator{}
	json.Unmarshal(aggrAsBytes, &a)

	bidkey = createbidkey(bidid)
	bAsBytes, err := stub.GetState(bidkey)
	if err != nil {
		return shim.Error(err.Error())
	} else if bAsBytes == nil {
		fmt.Println("Bid: " + bidid + " does not exist")
		return shim.Error("Bid: " + bidid + " does not exist")
	}

	b := BidProduce{}
	json.Unmarshal(bAsBytes, &b)
	
	if (b.Paidstatus == "PENDING" && b.Deliverystatus == "DELIVERED"){
		b.Paidstatus = "PAID"
		
			w.WalletBalance = w.WalletBalance - b.BidAmt
			a.WalletBalance = a.WalletBalance + b.BidAmt

			aggrDetAsBytes, _ := json.Marshal(a)
			err = stub.PutState(aggrkey, aggrDetAsBytes)
			if err != nil {
				return shim.Error(err.Error())
			}

			wDetAsBytes, _ := json.Marshal(w)
			err = stub.PutState(wholekey, wDetAsBytes)
			if err != nil {
				return shim.Error(err.Error())
			}

			bDetAsBytes, _ := json.Marshal(b)
			err = stub.PutState(bidkey, bDetAsBytes)
			if err != nil {
				return shim.Error(err.Error())
			}

			invkey := createinvkey(aggrid,strconv.FormatUint(b.Produceid,10))
			invAsBytes, err := stub.GetState(invkey)
			if err != nil {
					return shim.Error(err.Error())
			} else if invAsBytes == nil {
					fmt.Println("Inventory not found")
					return shim.Error("Inventory not found")
			}

			inv := Inventory{}
			json.Unmarshal(invAsBytes, &inv)
		
			inv.InvCount = inv.InvCount - b.Weight
			//fmt.Printf("inv.InvCount: %v" , inv.InvCount )

			invDetAsBytes, _ := json.Marshal(inv)
			err = stub.PutState(invkey, invDetAsBytes)
			if err != nil {
				return shim.Error(err.Error())
			}
		} else{
		return shim.Error("Bid: " + bidid + " is either pending for delivery or already PAID. Hence, Transaction Rejected!")
	}

	return shim.Success(nil)
}