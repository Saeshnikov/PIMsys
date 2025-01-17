import React, { useState, useEffect } from "react";
import { LogsClient } from "../../grpc/logs/logs_grpc_web_pb"; // Клиент gRPC
import { 
  GetLogsRequest, 
  GetLogsResponse, 
  Log
} from "../../grpc/logs/logs_pb"; // Сгенерированные сообщения
import { useNavigate } from "react-router-dom";
import { IconButton } from '@mui/material';
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


const columns = [
  { field: 'shopId', headerName: 'shopId', width: 120 },
  { field: 'branchId', headerName: 'branchId', width: 100 },
  { field: 'productId', headerName: 'productId', width: 120 },
  { field: 'info', headerName: 'info', width: 120 },
];

// Стили основной страницы
const StyledPage = styled(Box)(({ theme }) => ({
  display: "flex",
  flexDirection: "column",
  maxWidth: "800px",
  margin: "0 auto",
  padding: theme.spacing(4),
}));

const LogsPage = () => {
  const client = new LogsClient("http://localhost:8005"); // URL gRPC-сервера

  const [logs, setLogs] = useState([]);

  useEffect(() => {
    fetchLogs();
  }, []);

  const navigate = useNavigate(); // Хук для управления навигацией

  const token = localStorage.getItem("jwt_token");
  const metadata = {
    authorization: token,
  };

  const fetchLogs = async () => {
    const request = new GetLogsRequest();
    request.setProductId(1);
    client.getLogs(request, metadata, (err, response) => {
      if (err) {
        console.error("Ошибка загрузки списка логов:", err.message);
        return;
      }
      var i = 0;
      setLogs(response.getLogsList().map((log) => {
        var tmp = log.toObject();
        i= i+1;
        console.log(tmp);
        return {id: i, shopId: tmp.shopId, productId: tmp.productId, branchId: tmp.branchId, info: tmp.info};
      }));
    });
  };

  return (
    <div>
        <div>
        <StyledPage>
          <Divider sx={{ marginY: 3 }} />
          <Typography variant="h6" gutterBottom>
            Логи
          </Typography>
          <DataGrid rows={logs} columns={columns} pageSize={5} fullWidth/>
          <Divider sx={{ marginY: 3 }} />
        </StyledPage>
          <Typography variant="body2" color="textSecondary">
              (c) Brigada, Inc
          </Typography>
          </div>

    </div>

  );
};

export default LogsPage;
