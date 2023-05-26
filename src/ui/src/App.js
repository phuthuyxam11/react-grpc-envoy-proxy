import React, { useState, useEffect } from 'react';
import './App.css';

import {FileUploadRequest, MetaData} from "./protobuf/pb/service_upload_file_pb"
import {HrTimeSheetServiceClient} from "./protobuf/pb/service_upload_file_grpc_web_pb"

const client = new HrTimeSheetServiceClient("http://envoy:8080", null, null)


function App() {

  const [temp, setTemp] = useState(-9999);

  const getTemp = () => {
    console.log("called")

    let request = new FileUploadRequest()
    let metadata = new MetaData()
    metadata.setName("ddd")
    metadata.setType("sssss")
    request.setMetadata(metadata)

    console.log("dddd")
    let stream = client.uploadFile(request,{})
    console.log("dddd111")
    stream.on('data', function(response){
      console.log("sssss")
      setTemp(response.getValue())
    }).on("error", function (err) {
      console.log("errrr", err);
    });

  };

  useEffect(()=>{
    getTemp()
  },[]);


  return (
    <div className="App">
     eg
    </div>
  );
}

export default App;
