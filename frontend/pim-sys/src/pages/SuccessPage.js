import React from "react";
import Box from "@mui/material/Box";
import TextField from "@mui/material/TextField";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import { styled } from "@mui/material/styles";

const StyledBox = styled(Box)(({ theme }) => ({
  maxWidth: "400px",
  margin: "auto",
  padding: theme.spacing(4),
  textAlign: "center",
  backgroundColor: theme.palette.background.paper,
  borderRadius: theme.shape.borderRadius,
  boxShadow: theme.shadows[2],
  marginTop: theme.spacing(8),
}));

const SuccessPage = () => {
  const [companyData, setCompanyData] = React.useState({
    companyName: "",
    companyDescription: "",
  });

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setCompanyData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleSubmit = () => {
    alert(`Компания "${companyData.companyName}" добавлена успешно!`);
    // Здесь можно добавить логику для отправки данных компании на сервер.
  };

  return (
    <StyledBox>
      <Typography variant="h4" gutterBottom>
        Поздравляем! Вы успешно зарегистрировались
      </Typography>
      <Typography variant="body1" gutterBottom>
        Для того, чтобы начать работу с PIM-Системой, добавьте свою компанию
      </Typography>
      <TextField
        fullWidth
        margin="normal"
        label="Название компании"
        name="companyName"
        value={companyData.companyName}
        onChange={handleInputChange}
      />
      <TextField
        fullWidth
        margin="normal"
        label="Описание компании"
        name="companyDescription"
        value={companyData.companyDescription}
        onChange={handleInputChange}
        multiline
        rows={4}
      />
      <Button
        variant="contained"
        color="primary"
        fullWidth
        sx={{ mt: 2 }}
        onClick={handleSubmit}
      >
        Создать
      </Button>
    </StyledBox>
  );
};

export default SuccessPage;
