import React from "react";
import Box from "@mui/material/Box";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import Grid from "@mui/material/Grid";
import Alert from "@mui/material/Alert";
import { styled } from "@mui/material/styles";
import { RegisterRequest,LoginRequest } from "../../grpc/sso/sso_pb";
import { useNavigate } from "react-router-dom";


const StyledBox = styled(Box)(({ theme }) => ({
  width: "100%",
  maxWidth: "400px",
  padding: theme.spacing(4),
  borderRadius: theme.shape.borderRadius,
  backgroundColor: theme.palette.background.paper,
  boxShadow: theme.shadows[2],
}));

const RegisterForm = ({ authClient, onSwitchToLogin }) => {
  const [formData, setFormData] = React.useState({
    name: "",
    email: "",
    phone: "",
    password: "",
  });

  const [serverError, setServerError] = React.useState("");
  const navigate = useNavigate(); // Хук для управления навигацией

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevFormData) => ({
      ...prevFormData,
      [name]: value,
    }));
  };

  const register = () => {
    setServerError(""); // Очистка ошибки перед новым запросом

    const request = new RegisterRequest();
    request.setName(formData.name);
    request.setEmail(formData.email);
    request.setPassword(formData.password);
    request.setPhone(formData.phone)

    authClient.register(request, {}, (err) => {
      if (err) {
        setServerError(err.message || "Ошибка регистрации");
      } else {
        authClient.login(request, {}, (err, response) => {
          if (err) {
            setServerError(err.message || "Ошибка авторизации");
          } else {
            const jwtToken = response.getToken(); // Предположим, что токен приходит в ответе
            console.log(jwtToken)
            if (jwtToken) {
              // Сохраняем токен в localStorage
              localStorage.setItem("jwt_token", jwtToken);
    
              // Очистка ошибки при успешной авторизации
              setServerError(""); 
    
              // Перенаправление на страницу-пример
              navigate("/shop");
            } else {
              setServerError("Токен не получен, повторите попытку.");
            }
          }
        });
      }
    });
  };

  return (
    <StyledBox>
      <Typography variant="h5" gutterBottom>
        Начните управлять своими товарами
      </Typography>
      <Typography variant="body2" color="textSecondary" gutterBottom>
        Не упустите возможность повысить эффективность вашего бизнеса!
      </Typography>
      <Grid container spacing={2}>
        <Grid item xs={12}>
          <TextField
            fullWidth
            label="ФИО"
            name="name"
            value={formData.name}
            onChange={handleInputChange}
          />
        </Grid>
        <Grid item xs={12}>
          <TextField
            fullWidth
            label="Номер телефона"
            name="phone"
            value={formData.phone}
            onChange={handleInputChange}
          />
        </Grid>
        <Grid item xs={12}>
          <TextField
            fullWidth
            label="Email"
            name="email"
            value={formData.email}
            onChange={handleInputChange}
          />
        </Grid>
        <Grid item xs={12}>
          <TextField
            fullWidth
            label="Придумайте пароль"
            name="password"
            type="password"
            value={formData.password}
            onChange={handleInputChange}
          />
        </Grid>
      </Grid>
      {serverError && (
        <Box sx={{ mt: 2 }}>
          <Alert severity="error">{serverError}</Alert>
        </Box>
      )}
      <Box sx={{ mt: 2, display: "flex", justifyContent: "space-between" }}>
        <Button variant="contained" onClick={register}>
          Зарегистрироваться
        </Button>
        <Button variant="outlined" onClick={onSwitchToLogin}>
          Уже есть аккаунт
        </Button>
      </Box>
    </StyledBox>
  );
};

export default RegisterForm;
