\section{Project stuff}

<<data struct>>=
type HDat struct {
  UID    string `json:"uid"`
  time   string `json:"time"`
  e_code string `json:"ex-code"`
}
@


<<cc function manipulate data>>=

\section{Hyperledger/fabric stuff}
<<init function>>=
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
  return shim.Success(nil)
}
@
<<invoke function>>=
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

  function, args := APIstub.GetFunctionAndParameters()
  if function == "queryHDat" {
    return s.queryHDat(APIstub, args)
  } else if function == "createHDat" {
    return s.createHDat(APIstub, args)

  } else if function == "queryHDat" {
    return s.queryHDat(APIstub)
  }

  return shim.Error("Invalid Smart Contract function name.")
}
@


\section{drive shell script}


<<hdat.go>>=
<<copyright>>
package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "strconv"

  "github.com/hyperledger/fabric/core/chaincode/shim"
  sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

<<data struct>>

<<init function>>
<<invoke function>>

<<cc function manipulate data>>=
func main() {

  err := shim.Start(new(SmartContract))
  if err != nil {
    fmt.Printf("Error creating new Smart Contract: %s", err)
  }
}
@

func (s *SmartContract) createHDat(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

  if len(args) != 4 {
    return shim.Error("Incorrect number of arguments. Expecting 4")
  }

  var car = HDat{UID: args[1], time: args[2], e_code: args[3]}

  carAsBytes, _ := json.Marshal(car)
  APIstub.PutState(args[0], carAsBytes)

  return shim.Success(nil)
}
func (s *SmartContract) queryHDat(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

  if len(args) != 1 {
    return shim.Error("Incorrect number of arguments. Expecting 1")
  }

  carAsBytes, _ := APIstub.GetState(args[0])
  return shim.Success(carAsBytes)
}
@





<<copyright>>=
/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
@

