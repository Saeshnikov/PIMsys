import React, { useState, useEffect } from "react";
import { ShopClient } from "../../grpc/shop/shop_grpc_web_pb"; // Клиент gRPC
import { 
  ListShopsRequest, 
  NewShopRequest, 
  AlterShopRequest, 
  DeleteShopRequest, 
  ShopInfo 
} from "../../grpc/shop/shop_pb"; // Сгенерированные сообщения
import { useNavigate } from "react-router-dom";
import { IconButton } from '@mui/material';
import DeleteForeverIcon from '@mui/icons-material/DeleteForever';
import EditIcon from '@mui/icons-material/Edit';
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

// Стили основной страницы
const StyledPage = styled(Box)(({ theme }) => ({
  display: "flex",
  flexDirection: "column",
  maxWidth: "800px",
  margin: "0 auto",
  padding: theme.spacing(4),
}));

// Стили для формы добавления филиалов
const StyledForm = styled(Box)(({ theme }) => ({
  display: "flex",
  flexDirection: "column",
  gap: theme.spacing(2),
  backgroundColor: theme.palette.background.paper,
  padding: theme.spacing(4),
  borderRadius: theme.shape.borderRadius,
  boxShadow: theme.shadows[1],
}));

// Стили для карточек магазинов
const StyledCard = styled(Card)(({ theme }) => ({
  padding: theme.spacing(2),
  boxShadow: theme.shadows[1],
  borderRadius: theme.shape.borderRadius,
}));

const ShopPage = () => {
  const deleteIcon = <DeleteForeverIcon />;
  const editIcon = <EditIcon />;
  const client = new ShopClient("http://localhost:8001"); // URL gRPC-сервера

  const [shops, setShops] = useState([]);
  const [newShop, setNewShop] = useState({ name: "", description: "", url: "" });
  const [editShop, setEditShop] = useState(null);

  useEffect(() => {
    fetchShops();
  }, []);

  const navigate = useNavigate(); // Хук для управления навигацией

  const token = localStorage.getItem("jwt_token");
  const metadata = {
    authorization: token,
  };

  const fetchShops = async () => {
    const request = new ListShopsRequest();
    client.listShops(request, metadata, (err, response) => {
      if (err) {
        console.error("Ошибка загрузки списка магазинов:", err.message);
        return;
      }
      setShops(response.getInfoList().map((shop) => shop.toObject()));
    });
  };

  const handleAddShop = async () => {
    const request = new NewShopRequest();
    request.setName(newShop.name);
    request.setDescription(newShop.description);
    request.setUrl(newShop.url);

    client.newShop(request, metadata, (err, response) => {
      if (err) {
        console.error("Ошибка добавления магазина:", err.message);
        alert("Ошибка добавления магазина:", err.message)
        return;
      }
      setNewShop({ name: "", description: "", url: "" });
      fetchShops();
    });
  };

  const handleEditShop = async () => {
    const request = new AlterShopRequest();
    const shopInfo = new ShopInfo();
    shopInfo.setShopId(editShop.shopId);
    shopInfo.setName(editShop.name);
    shopInfo.setDescription(editShop.description);
    shopInfo.setUrl(editShop.url);

    request.setShopId(editShop.shopId);
    request.setShopInfo(shopInfo);
    

    client.alterShop(request, metadata, (err, response) => {
      if (err) {
        console.error("Ошибка изменения магазина:", err.message);
        alert("Ошибка изменения магазина:", err.message)
        return;
      }
      setEditShop(null);
      fetchShops();
    });
  };

  const handleDeleteShop = async (shop_id) => {
    const request = new DeleteShopRequest();
    request.setShopId(shop_id);

    client.deleteShop(request, metadata, (err, response) => {
      if (err) {
        console.error("Ошибка удаления магазина:", err.message);
        alert("Ошибка удаления магазина:", err.message)
        return;
      }
      fetchShops();
    });
  };

  const handleOpenBranches = async (shop) => {
    const request = new ShopInfo();
    localStorage.setItem("shop_name", shop.name);
    localStorage.setItem("shop_description", shop.description);

    navigate(`/shop/${shop.shopId}/branches`);
  };

  return (
    <div>
        <div>
        <Button
          variant="contained"
          color="primary"
          onClick={() => {navigate(`/logs`)}}
          sx={{ alignSelf: "flex-start" }}
        >
          Логи
        </Button>
        <Button
          variant="contained"
          color="primary"
          onClick={() => {navigate(`/graphs`)}}
          sx={{ alignSelf: "flex-start" }}
        >
          Графики
        </Button>
        <StyledPage>
          <Divider sx={{ marginY: 3 }} />
          <Typography variant="h6" gutterBottom>
            Компании
          </Typography>
          <StyledForm>
            <Typography variant="subtitle1">Добавить новую компанию</Typography>
            <TextField
              label="Название компании"
              fullWidth
              value={newShop.name}
              onChange={(e) => setNewShop({ ...newShop, name: e.target.value })}
            />
            <TextField
              label="Описание компании"
              fullWidth
              value={newShop.description}
              onChange={(e) => setNewShop({ ...newShop, description: e.target.value })}
              multiline
              rows={3}
            />
            <TextField
              label="URL"
              fullWidth
              value={newShop.url}
              onChange={(e) => setNewShop({ ...newShop, url: e.target.value })}
            />
            <Button
              variant="contained"
              color="primary"
              onClick={handleAddShop}
              sx={{ alignSelf: "flex-start" }}
            >
              Добавить компанию
            </Button>
          </StyledForm>
          <Divider sx={{ marginY: 3 }} />
          <Grid container spacing={2}>
            {shops.map((shop) => (
              <Grid item xs={12} sm={6} md={4} key={shop.shopId}>
                  {(editShop==null || editShop.shopId!==shop.shopId) && (<StyledCard>
                    <Grid container spacing={2}>
                      <Grid item xs={7}>
                        <Typography variant="subtitle1" gutterBottom>
                          {shop.name}
                        
                        </Typography>
                      </Grid>
                      <Grid item xs={2}>
                        <IconButton onClick={() => setEditShop(shop)}
                        variant="filled"
                        color="primary"
                        sx={{ alignSelf: "flex-end" }}
                        >{editIcon}</IconButton>
                      </Grid>
                      <Grid item xs={2}>
                        <IconButton onClick={() => handleDeleteShop(shop.shopId)}
                        variant="filled"
                        color="danger"
                        sx={{ alignSelf: "flex-end" }}
                        >{deleteIcon}</IconButton>
                      </Grid>
                    </Grid>
                    
                    
                    
                    <Chip
                      label="Филиалы"
                      color="primary"
                      variant="outlined"
                      sx={{ marginBottom: 1 }}
                      onClick={() => handleOpenBranches(shop)}
                    />
                    <Typography variant="body2" color="textSecondary">
                      {shop.description}
                    </Typography>
                  </StyledCard>
                  )}
                  {editShop!=null && editShop.shopId===shop.shopId && (
                  <div>
                    <StyledCard>
                        <TextField
                          label="Название компании"
                          fullWidth
                          placeholder="Название"
                          value={editShop.name}
                          onChange={(e) => setEditShop({ ...editShop, name: e.target.value })}
                        />
                        <TextField
                        label="Описание компании"
                        fullWidth
                        placeholder="Описание"
                        value={editShop.description}
                        onChange={(e) => setEditShop({ ...editShop, description: e.target.value })}
                        />
                        <TextField
                          label="URL"
                          fullWidth
                          placeholder="URL"
                          value={editShop.url}
                          onChange={(e) => setEditShop({ ...editShop, url: e.target.value })}
                      />
                      
                      <Button onClick={handleEditShop} 
                      variant="contained"
                      color="primary"
                      sx={{ alignSelf: "flex-start" }}
                      >Сохранить</Button>

                      <Button onClick={() => setEditShop(null)}
                      variant="contained"
                      color="primary"
                      sx={{ alignSelf: "flex-start" }
                      }>Отмена</Button>
                    </StyledCard>
                    
                    
                  </div>
                )}
              </Grid>
            ))}
          </Grid>
            
            
          </StyledPage>
          <Typography variant="body2" color="textSecondary">
              (c) Brigada, Inc
          </Typography>
          </div>

    </div>

  );
};

export default ShopPage;
