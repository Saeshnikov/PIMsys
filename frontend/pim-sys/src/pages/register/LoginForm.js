import React from "react";
import Box from "@mui/material/Box";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import Grid from "@mui/material/Grid";
import Alert from "@mui/material/Alert";
import { styled } from "@mui/material/styles";
import { LoginRequest } from "../../grpc/sso/sso_pb";
import { useNavigate } from "react-router-dom";

const StyledBox = styled(Box)(({ theme }) => ({
  display: "flex",
  flexDirection: "column",
  alignItems: "center",
  justifyContent: "center",
  maxWidth: "400px",
  margin: "0 auto",
  padding: theme.spacing(4),
  borderRadius: theme.shape.borderRadius,
  boxShadow: theme.shadows[3],
  backgroundColor: theme.palette.background.paper,
}));

const LoginForm = ({ authClient, onSwitchToRegister }) => {
  const [formData, setFormData] = React.useState({
    email: "",
    password: "",
  });

  const [serverError, setServerError] = React.useState("");
  const navigate = useNavigate(); // Для навигации после успешного логина

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevFormData) => ({
      ...prevFormData,
      [name]: value,
    }));
  };

  const login = () => {
    setServerError(""); // Очистка ошибки перед новым запросом

    const request = new LoginRequest();
    request.setEmail(formData.email);
    request.setPassword(formData.password);

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
  };

  return (
    <StyledBox>
      <Typography variant="h5" gutterBottom>
        Авторизация
      </Typography>
      <Grid container spacing={2}>
        <Grid item xs={12}>
          <TextField
            fullWidth
            required
            label="Email"
            name="email"
            value={formData.email}
            onChange={handleInputChange}
            helperText=""
          />
        </Grid>
        <Grid item xs={12}>
          <TextField
            fullWidth
            required
            label="Пароль"
            name="password"
            type="password"
            value={formData.password}
            onChange={handleInputChange}
            helperText=""
          />
        </Grid>
      </Grid>
      {serverError && (
        <Box sx={{ mt: 2, width: "100%" }}>
          <Alert severity="error">{serverError}</Alert>
        </Box>
      )}
      <Box sx={{ mt: 2, display: "flex", gap: 1 }}>
        <Button variant="contained" onClick={login}>
          Войти
        </Button>
        <Button variant="outlined" onClick={onSwitchToRegister}>
          Назад к регистрации
        </Button>
      </Box>
    </StyledBox>
  );
};

export default LoginForm;
