import React from "react";
import Box from '@mui/material/Box';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';


import { RegisterRequest } from "./grpc/sso/sso_pb";
import { AuthClient } from "./grpc/sso/sso_grpc_web_pb";

let auth_client = new AuthClient("http://localhost:8000");


function App() {

  const [formData, setFormData] = React.useState({
    name: "",
    email: "",
    phone: "",
    password: "",
  });

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevFormData) => ({
      ...prevFormData,
      [name]: value,
    }));
  };


  function register(){
    auth_client.register(new RegisterRequest(formData.name,formData.email,formData.password), {}, (err, response) => {
      if (err) {
        console.log(err)
      }else{
        console.log(response)
      }
    });
  }

  return (
    <div className="component-app">
        <Box
    component="form"
    sx={{ '& .MuiTextField-root': { m: 1, width: '25ch' } }}
    noValidate
    autoComplete="off"
  >


  <div>
    <TextField
      required
      id="outlined-required"
      label="FIO"
      name = "name"
      defaultValue="Egor"
      onChange={handleInputChange}
      value={formData.name}
    />
    <TextField
      required
      id="outlined-required"
      label="EMAIL"
      name="email"
      defaultValue="Egor@egor"
      onChange={handleInputChange}
      value={formData.email}
    />
    <TextField
      required
      id="outlined"
      label="PHONE"
      name="phone"
      defaultValue="1234"
      onChange={handleInputChange}
      value={formData.phone}
    />
    <TextField
      id="outlined-password-input"
      label="Password"
      name="password"
      type="password"
      autoComplete="current-password"
      onChange={handleInputChange}
      value={formData.password}
    />
    </div>  
    <Button variant="contained" onClick={register}>register</Button>
    <Button variant="contained">have account</Button>
  </Box>
    </div>
  );

}

export default App;