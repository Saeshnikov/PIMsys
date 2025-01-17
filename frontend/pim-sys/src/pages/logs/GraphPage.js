import React, { useState, useEffect } from "react";
import { LogsClient } from "../../grpc/logs/logs_grpc_web_pb"; // Клиент gRPC
import { 
  GetGraphRequest, 
  GetGraphResponse, 
  Graph
} from "../../grpc/logs/logs_pb"; // Сгенерированные сообщения
import { useNavigate } from "react-router-dom";
import { IconButton } from '@mui/material';
import { DateField } from '@mui/x-date-pickers/DateField';
import {
  Button,
  TextField,
  Box,
  Typography,
  Grid,
  RadioGroup,
  FormControlLabel,
  Radio,
  Divider,
  Card,
  CardContent,
  Chip,
} from "@mui/material";
import { styled } from "@mui/material/styles";
import { DataGrid } from '@mui/x-data-grid';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import dayjs from "dayjs"; // Для работы с датами
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from "recharts";


// Стили основной страницы
const StyledPage = styled(Box)(({ theme }) => ({
  display: "flex",
  flexDirection: "column",
  maxWidth: "800px",
  margin: "0 auto",
  padding: theme.spacing(4),
}));

const GraphsPage = () => {
  const client = new LogsClient("http://localhost:8005"); // URL gRPC-сервера

  const [data, setData] = useState([]);
  const [interval, setInterval] = useState(0);
  const [dateFrom, setDateFrom] = useState(dayjs());
  const [dateTo, setDateTo] = useState(dayjs());

  useEffect(() => {
    fetchData();
  }, []);

  const navigate = useNavigate(); // Хук для управления навигацией

  const token = localStorage.getItem("jwt_token");
  const metadata = {
    authorization: token,
  };

  const fetchData = async () => {
    const request = new GetGraphRequest();
    request.setInterval(interval);
    console.log(dateFrom.unix())
    request.setDateFrom(dateFrom.unix());
    request.setDateTo(dateTo.unix());
    request.setProductId(1);
    client.getGraph(request, metadata, (err, response) => {
      if (err) {
        console.error("Ошибка загрузки списка данных:", err.message);
        return;
      }
      setData(response.getGraphsList().map((log) => {
        const microData = log.toObject();
        console.log(new Date(microData.date* 1000));
        return {
          ...microData,
          date: new Date(microData.date* 1000), // Преобразование timestamp в строку даты
        };
      }));
    });
  };

  return (
    <div>
        <div>
        <StyledPage>
          <Divider sx={{ marginY: 3 }} />
          <Typography variant="h6" gutterBottom>
            Графики
          </Typography>
          <RadioGroup
            row
            value={interval}
            onChange={(e) => setInterval(e.target.value)}
          >
            <FormControlLabel
              value={0}
              control={<Radio />}
              label="Дни"
            />
            <FormControlLabel
              value={1}
              control={<Radio />}
              label="Месяцы"
            />
          </RadioGroup>
          <LocalizationProvider dateAdapter={AdapterDayjs}>
            <DateField
              label="Date From"
              value={dateFrom}
              onChange={(e) => setDateFrom(e)}
            />
          </LocalizationProvider>
          <LocalizationProvider dateAdapter={AdapterDayjs}>
            <DateField
              label="Date To"
              value={dateTo}
              onChange={(e) =>setDateTo(e)}
            />
          </LocalizationProvider>
          <Button
            variant="contained"
            color="primary"
            onClick={() => {fetchData()}}
            sx={{ alignSelf: "flex-start" }}
          >
            Запрос данных
          </Button>
          <Divider sx={{ marginY: 3 }} />
          <LineChart
            width={800}
            height={400}
            data={data}
            margin={{ top: 20, right: 30, left: 20, bottom: 5 }}
          >
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis dataKey="date" />
            <YAxis />
            <Tooltip />
            <Legend />
            <Line type="monotone" dataKey="totalSales" stroke="#8884d8" name="Total Sales" />
          </LineChart>
          <LineChart
            width={800}
            height={400}
            data={data}
            margin={{ top: 20, right: 30, left: 20, bottom: 5 }}
          >
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis dataKey="date" />
            <YAxis />
            <Tooltip />
            <Legend />
            <Line type="monotone" dataKey="totalQuantity" stroke="#82ca9d" name="Total Quantity" />
          </LineChart>
        </StyledPage>
          <Typography variant="body2" color="textSecondary">
              (c) Brigada, Inc
          </Typography>
          </div>

    </div>

  );
};

export default GraphsPage;
