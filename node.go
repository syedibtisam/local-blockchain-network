package main

import (
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

var Neighbours []string
var Transactions_list []string

// /////////////////////////////////////////////////
// Block Chain
type Block struct {
	// data                   string
	MerkleTree             *Node
	Nonce                  string
	Previous_block_address *Block
	Previous_block_hash    string
	Block_number           int
	Block_hash             string
}

type Blockchain struct {
	genesis_block *Block
	last_block    *Block
}

func New_Block(merkelTree *Node, nonce string) *Block {
	new_block := new(Block)
	new_block.MerkleTree = merkelTree
	new_block.Nonce = nonce

	return new_block
}

func Mine_Block(new_block *Block) bool {
	number := 0
	threshold := 1000
	// c := make(chan int)
	for {
		if threshold < number {
			return false
		}
		if strconv.Itoa(number) == new_block.Nonce {
			return true
		} else {
			// go Check_block_exist(new_block.Block_number, c)
			// number := <-c
			// if number == 20 {
			// 	return false
			// }
			number++
		}
	}
}

func Add_To_Blockchain(previous_block_address *Block, new_block *Block) *Block {

	if Mine_Block(new_block) {
		if previous_block_address == nil {
			new_block.Previous_block_address = nil
			new_block.Previous_block_hash = ""
			new_block.Block_number = 1

		} else {
			new_block.Previous_block_address = previous_block_address
			new_block.Previous_block_hash = previous_block_address.Block_hash
			new_block.Block_number = previous_block_address.Block_number + 1

		}
		new_block.Block_hash = Calculate_Hash(new_block)
		return new_block
	}
	return previous_block_address
}

func Display_Blocks(last_block *Block) {
	list := last_block
	if list == nil {
		println("No blocks in the Blockchain")
	}
	for list != nil {
		fmt.Println("-------------------- Block", list.Block_number, "--------------------")
		fmt.Println("Transaction (Merkle Tree) : ")
		DisplayMerkelTree(list.MerkleTree)
		fmt.Println("Nonce                     : ", list.Nonce)
		fmt.Println("Block Hash                : ", list.Block_hash)
		fmt.Println("Previous block Hash       : ", list.Previous_block_hash)
		fmt.Println("")
		list = list.Previous_block_address
	}
	println("")
	println("Blockchain end")
}

func Create_Blockchain() *Blockchain {
	temp_block_chain := new(Blockchain)
	temp_block_chain.genesis_block = nil
	temp_block_chain.last_block = nil
	return temp_block_chain
}
func Calculate_Hash(block *Block) string {
	str2 := strconv.Itoa(block.Block_number)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(block.MerkleTree.hash+block.Nonce+str2+block.Previous_block_hash)))
}

func Verify_Chain(last_block *Block) {
	list := last_block

	if list == nil {
		println("No blocks in the Blockchain")
	}
	for list != nil {
		if list.Previous_block_address != nil {
			if list.Previous_block_hash == Calculate_Hash(list.Previous_block_address) {
				list = list.Previous_block_address
			} else {
				fmt.Println("Block", list.Block_number-1, "is tempered")
				break
			}
		} else {
			fmt.Println("All blocks are verified and not tempered")
			break
		}
	}
	println("")

}

func Change_Block(block *Block) {
	Change_Block_Menu()
	user_input := ""
	fmt.Print("Enter number : ")
	fmt.Scanln(&user_input)

	if user_input == "1" {
		fmt.Print("Enter new Block number : ")
		fmt.Scanln(&user_input)
		Int, err := strconv.Atoi(user_input)
		if err == nil {
			block.Block_number = Int
		} else {
			fmt.Println("Should be integer")
		}

	} else if user_input == "2" {
		fmt.Print("Enter Nonce: ")
		fmt.Scanln(&user_input)
		block.Nonce = user_input

	} else if user_input == "3" {
		fmt.Print("Enter New Hash: ")
		fmt.Scanln(&user_input)
		block.Previous_block_hash = user_input
	} //else if user_input == "4" {
	// 	fmt.Print("Enter New Data: ")
	// 	fmt.Scanln(&user_input)
	// 	block.data = user_input
	// }

}

func Menu() {
	println("-----------------------------")
	println("1) Display Blocks")
	println("2) Add new block")
	println("6) Check node list of transactions")
	println("7) Check current neighbours")
	println("9) Exit")
	println("-----------------------------")

}

func Change_Block_Menu() {
	fmt.Println("1) Change Block Number")
	fmt.Println("2) Change Nonce")
	fmt.Println("3) Change Previous Block Hash")
	fmt.Println("4) Change Data")
}

// //////////////////////////////////////////////////////////////////
// Merkel Tree
type Node struct {
	hash  string
	left  *Node
	right *Node
}

func getLeft(n *Node) *Node {
	return n.left
}
func setLeft(n *Node, x *Node) {
	n.left = x
}
func getRight(n *Node) *Node {
	return n.right
}
func setRight(n *Node, x *Node) {
	n.right = x
}
func getHash(n *Node) string {
	return n.hash
}
func setHash(n *Node, x string) {
	n.hash = x
}

func generateTree(dataBlocks []string) *Node {

	var arr1 = make([]*Node, len(dataBlocks))

	for i := 0; i < len(dataBlocks); i++ {

		nodeObj := new(Node)
		setLeft(nodeObj, nil)
		setRight(nodeObj, nil)
		setHash(nodeObj, fmt.Sprintf("%x", sha256.Sum256([]byte(dataBlocks[i]))))
		arr1[i] = nodeObj
	}

	return buildTree(arr1)
}
func buildTree(children []*Node) *Node {

	var parents = make([]*Node, len(children))

	for len(children) != 1 {
		var index = 0
		var length = len(children)
		var i = 0
		for index < length {
			leftChild := children[index]
			rightChild := new(Node)

			if (index + 1) < length {
				rightChild = children[index+1]
			} else {
				nodeObj := new(Node)
				setLeft(nodeObj, nil)
				setRight(nodeObj, nil)
				setHash(nodeObj, getHash(leftChild))
				rightChild = nodeObj
			}
			var parentHash = fmt.Sprintf("%x", sha256.Sum256([]byte(getHash(leftChild)+getHash(rightChild))))
			nodeObj := new(Node)
			setLeft(nodeObj, leftChild)
			setRight(nodeObj, rightChild)
			setHash(nodeObj, parentHash)

			parents[i] = nodeObj
			i++
			index += 2
		}
		children = parents[0:i]

		parents = parents[0:0]
		parents = parents[0:len(children)]

	}
	return children[0]
}
func DisplayMerkelTree(root *Node) {
	if root == nil {
		return
	}

	if getLeft(root) == nil && getRight(root) == nil {
		fmt.Println(getHash(root))
	}
	queue := make([]*Node, 0)
	// Push queue
	queue = append(queue, root)
	queue = append(queue, nil)

	for !(len(queue) == 0) {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			fmt.Println(getHash(node))
		} else {
			fmt.Println()
			if !(len(queue) == 0) {
				queue = append(queue, nil)
			}
		}

		if node != nil && getLeft(node) != nil {
			queue = append(queue, getLeft(node))
		}

		if node != nil && getRight(node) != nil {
			queue = append(queue, getRight(node))
		}

	}

}

func Get_Transactions(transactions []string) *Node {
	len := 5
	var dataBlocks = make([]string, len)
	dataBlocks[0] = transactions[0]
	dataBlocks[1] = transactions[1]
	dataBlocks[2] = transactions[2]
	dataBlocks[3] = transactions[3]
	dataBlocks[4] = transactions[4]
	return generateTree(dataBlocks)
	// DisplayMerkelTree(root)
}
func Generate_block(transactions []string, blocknumber int, nonce string) TempBlock {
	var send_block TempBlock
	for i := 0; i < len(transactions); i++ {
		send_block.Transactions_node = append(send_block.Transactions_node, transactions[i])
	}
	send_block.BlockNumber = blocknumber
	send_block.nonce = nonce
	return send_block
}

type TempBlock struct {
	Transactions_node []string
	BlockNumber       int
	nonce             string
}

// /////////////////////////////////////////////////
const (
	SERVER_HOST    = "localhost"
	Bootstrap_PORT = "8000"
	SERVER_TYPE    = "tcp"
)

//	func update_neighbours() {
//		fmt.Println("Node making connection to bootstrap node to update neighbours")
//		connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+Bootstrap_PORT)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println("Connected to bootstrap")
//		///send some data
//		_, err = connection.Write([]byte("get_total_ports")) // sending data to server
//		buffer := make([]byte, 1024)
//		mLen, err := connection.Read(buffer) // reading from server
//		if err != nil {
//			fmt.Println("Error reading:", err.Error())
//		}
//		// fmt.Println("Received: ", string(buffer[:mLen]))
//		connection.Close()
//		updated_biggest_port := (string(buffer[:mLen]))
//		fmt.Println(updated_biggest_port)
//		// neighbours = append(neighbours, updated_biggest_port)
//	}
var block_chain = Create_Blockchain()

func main() {
	// Transactions_list = append(Transactions_list, "temp1")
	// Transactions_list = append(Transactions_list, "temp2")
	// Transactions_list = append(Transactions_list, "temp3")
	// Transactions_list = append(Transactions_list, "temp4")
	//establish connection to bootstrap node
	//to get port number where the node will be listening
	//to get neighbours to make connections with some other nodes in the network
	fmt.Println("Node making connection to bootstrap node to get the port number")
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+Bootstrap_PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("Connected to bootstrap")
	///send some data
	_, err = connection.Write([]byte("get_port")) // sending data to server
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer) // reading from server
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// fmt.Println("Received: ", string(buffer[:mLen]))
	connection.Close()
	node_server_port := (string(buffer[:mLen]))

	go BecomeServer(node_server_port)
	system_alive := true
	user_input := ""
	for system_alive {
		Menu()

		fmt.Print("Choice : ")
		fmt.Scanln(&user_input)
		fmt.Println("")

		if user_input == "9" {
			system_alive = false
		} else if user_input == "1" {
			Display_Blocks(block_chain.last_block)
		} else if user_input == "2" {
			fmt.Print("Enter Transaction: ")
			fmt.Scanln(&user_input)
			Send_transaction_to_all_neighbours(user_input, node_server_port) //sender port

		} else if user_input == "6" {
			fmt.Println("transactions : ")
			for _, ele := range Transactions_list { // you can escape index by _ keyword
				fmt.Println(ele)
			}
		} else if user_input == "7" {
			fmt.Println("neighbours : ")
			for _, ele := range Neighbours { // you can escape index by _ keyword
				fmt.Println(ele)
			}
		}

		user_input = ""
		println("")
	}
}

type CommandType struct {
	Command     string
	Client_port string
	// PrevPointer *Block
}

func Flood_across_Network(transaction_to_flood string, node_server_port string) {
	length := len(Neighbours)
	for i := 0; i < length; i++ {
		if Neighbours[i] == "" {
			continue
		}
		Send_transaction_to_node(transaction_to_flood, Neighbours[i], node_server_port)
	}
}
func Send_transaction_to_all_neighbours(transaction string, node_server_port string) {
	Transactions_list = append(Transactions_list, transaction)
	//flood across network
	Flood_across_Network(transaction, node_server_port)
	//mine if length crosses limit 5
	if len(Transactions_list) == 5 {
		block_to_add := New_Block(Get_Transactions(Transactions_list), strconv.Itoa(rand.Intn(1000)))
		Mine_and_flood_across_network(block_to_add, Transactions_list)
		block_chain.last_block = Add_To_Blockchain(block_chain.last_block, block_to_add)
		// fmt.Println("5 transactions done")
		// create new block
		//mine it
		//flood across network
		//pop out transactions that are being mined and verified
	}
}
func Send_transaction_to_node(transaction_to_flood string, node_port string, node_server_port string) {
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+node_port)
	if err != nil {
		fmt.Println("Error in send transaction node: ", err)
		return
	}
	transaction_type := &CommandType{"transaction", node_server_port}
	gobEncoder := gob.NewEncoder(connection)
	errr := gobEncoder.Encode(transaction_type)
	if errr != nil {
		log.Println(errr)
	}
	transaction := &CommandType{transaction_to_flood, node_server_port} // transaction, port of sender
	errrr := gobEncoder.Encode(transaction)
	if err != nil {
		log.Println(errrr)
	}
	connection.Close()
}
func Mine_and_flood_across_network(new_block *Block, transactions []string) {

	if Mine_Block(new_block) { //mining
		fmt.Println("mining done")
		length := len(Neighbours)
		for i := 0; i < length; i++ {
			if Neighbours[i] == "" {
				continue
			}
			Send_one_mined_block_to_node(new_block, Neighbours[i], transactions)
		}
	}
}
func Check_block_exist(blockNumber int, number chan int) bool {
	list := block_chain.last_block
	if list == nil {
		number <- 10
		return true
	} else {
		for list != nil {
			if blockNumber == list.Block_number {
				number <- 20 //not add
				return false
			}
			list = list.Previous_block_address
		}
	}
	number <- 10 //to add
	return true
}
func Send_one_mined_block_to_node(new_block *Block, node_port string, transactions []string) {
	fmt.Println("Sending mined block " + new_block.Nonce + " to port " + node_port)
	connection, erri := net.Dial(SERVER_TYPE, SERVER_HOST+":"+node_port)
	if erri != nil {
		fmt.Println("Error in send transaction node: ", erri)
		return
	}
	transaction_type := &CommandType{"mined_block", node_port}
	gobEncoder := gob.NewEncoder(connection)
	err := gobEncoder.Encode(transaction_type)
	if err != nil {
		log.Println(err)
	}
	send_block := Generate_block(transactions, new_block.Block_number, new_block.Nonce)
	block := &send_block
	errr := gobEncoder.Encode(block)
	if errr != nil {
		log.Println(errr)
	}

}
func Display_single_block(block TempBlock) {
	fmt.Println("-------------------- Block", block.BlockNumber, "--------------------")
	fmt.Println("Transactions")
	for i := 0; i < len(block.Transactions_node); i++ {
		fmt.Println(block.Transactions_node[i])
	}
	fmt.Println("--------------------")
}
func ProcessClient(connection net.Conn, port_to_listen string) {
	var recvdBlock CommandType
	dec := gob.NewDecoder(connection)
	err := dec.Decode(&recvdBlock)
	if err != nil {
		fmt.Println("Error in processing client ", err)
	}

	/////////////////////////////////////////////////////
	// when the client sends transaction

	if recvdBlock.Command == "transaction" {

		var transaction CommandType
		errr := dec.Decode(&transaction)
		if errr != nil {
			//handle error
			fmt.Println("Error after taking transaction ", errr)
		}

		to_add := true
		for _, ele := range Neighbours {
			if ele == transaction.Client_port {
				to_add = false
			}
		}
		if to_add {
			// fmt.Println("Port added : ", transaction.Client_port)
			Neighbours = append(Neighbours, transaction.Client_port)
		} else {
			// fmt.Println("connected port not added")
		}
		////////////////////////////////////////////////////
		// adding transactions to list
		add_transaction := true
		for _, ele := range Transactions_list {
			if ele == transaction.Command {
				add_transaction = false
				break
			}
			// fmt.Println(ele)
		}
		if add_transaction {
			//1 add to list of transaction
			Transactions_list = append(Transactions_list, transaction.Command)
			//flood across network
			go Flood_across_Network(transaction.Command, port_to_listen)
			//mine if length crosses limit 5
			if len(Transactions_list) == 5 {
				fmt.Println("5 transactions done")
				// create new block
				block_to_add := New_Block(Get_Transactions(Transactions_list), strconv.Itoa(rand.Intn(1000)))
				Mine_and_flood_across_network(block_to_add, Transactions_list)
				block_chain.last_block = Add_To_Blockchain(block_chain.last_block, block_to_add)
				//mine it
				//flood across network
				//pop out transactions that are being mined and verified
			}
		}
	} else if recvdBlock.Command == "mined_block" {
		fmt.Println("receiving mined block")
		var mined_block TempBlock

		errr := dec.Decode(&mined_block)
		if errr != nil {
			//handle error
			fmt.Println("Error after taking transaction ", errr)
		}
		c := make(chan int)
		if Check_block_exist(mined_block.BlockNumber, c) {
			Display_single_block(mined_block)
			fmt.Println("Adding to local blockchain")
			block_to_add := New_Block(Get_Transactions(mined_block.Transactions_node), mined_block.nonce)
			block_chain.last_block = Add_To_Blockchain(block_chain.last_block, block_to_add)
		} else {
			fmt.Println("Already in blockchain")
		}
	}
	connection.Close()
}

func Random_neighbours(max_port int) {

	rand.Seed(time.Now().UnixNano())
	min := 1000
	max_n := max_port - 1000
	max_neighbours := max_n / 3
	to_add := true
	for i := 0; i < max_neighbours; i++ {
		length := len(Neighbours)
		random_port := rand.Intn(max_port-min+1) + min
		for j := 0; j < length; j++ {
			if strconv.Itoa(random_port) == Neighbours[j] {
				to_add = false
				break
			}
		}
		if to_add {
			Neighbours = append(Neighbours, strconv.Itoa(random_port))
		} else {
			to_add = true
		}
	}
}
func BecomeServer(port_to_listen string) {
	// fmt.Println("Received port number and neighbours")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+port_to_listen) // listening for clients
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// adding neighbours
	i, err := strconv.Atoi(port_to_listen)

	if i > 1000 {
		if i < 1006 { // add all neighbours
			for j := 1000; j < i; j++ {
				Neighbours = append(Neighbours, strconv.Itoa(j))
			}
		} else { // add random neighbours
			Random_neighbours(i)
		}
	}
	defer server.Close()
	fmt.Println("Listening on " + SERVER_HOST + ":" + port_to_listen)
	// fmt.Println("Waiting for client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		// fmt.Println("client connected")
		go ProcessClient(connection, port_to_listen)
	}
}
