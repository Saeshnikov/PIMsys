import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
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
import { ShopClient } from "../../grpc/shop/shop_grpc_web_pb";
import {
  ListShopsRequest,
  NewShopRequest,
  DeleteShopRequest,
} from "../../grpc/shop/shop_pb";
import { useNavigate } from "react-router-dom";

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
  const navigate = useNavigate();
  const [shops, setShops] = useState([]);
  const [newShopName, setNewShopName] = useState("");
  const [newShopDescription, setNewShopDescription] = useState("");
  const [newShopUrl, setNewShopUrl] = useState("");
  const [shopType, setShopType] = useState("online");

  const shopClient = new ShopClient("http://localhost:8001", null, null);

  const loadShops = () => {
    const request = new ListShopsRequest();
    const token = localStorage.getItem("jwt_token");
    const metadata = {
      authorization: token,
    };
    shopClient.listShops(request, metadata, (err, response) => {
      if (err) {
        console.error("Error loading shops", err);
      } else {
        setShops(response.getInfoList());
      }
    });
  };

  const createNewShop = () => {
    const request = new NewShopRequest();
    request.setName(newShopName);
    request.setDescription(newShopDescription);
    request.setUrl(newShopUrl);
    request.setType(shopType);

    const token = localStorage.getItem("jwt_token");
    const metadata = {
      authorization: token,
    };

    shopClient.newShop(request, metadata, (err, response) => {
      if (err) {
        console.error("Error creating shop", err);
      } else {
        loadShops();
        setNewShopName("");
        setNewShopDescription("");
        setNewShopUrl("");
        setShopType("online");
      }
    });
  };

  const deleteShop = (shopId) => {
    const request = new DeleteShopRequest();
    request.setShopId(shopId);

    const token = localStorage.getItem("jwt_token");
    const metadata = {
      authorization: token,
    };

    shopClient.deleteShop(request, metadata, (err, response) => {
      if (err) {
        console.error("Error deleting shop", err);
      } else {
        loadShops();
      }
    });
  };

  useEffect(() => {
    const token = localStorage.getItem("jwt_token");

    if (!token) {
      navigate("/");
      return;
    }
    loadShops();
  }, []);

  return (
    <StyledPage>
      <Typography variant="h4" gutterBottom>
        Компания
      </Typography>
      <Typography variant="h5" gutterBottom>
        RAGGIES
      </Typography>
      <Typography variant="body1" gutterBottom>
        Мы предлагаем вам эффективное решение для управления товарами в магазинах.
        Наша система поможет вам оптимизировать процессы, связанные с ассортиментом,
        и улучшить качество обслуживания клиентов.
      </Typography>
      <Divider sx={{ marginY: 3 }} />
      <Typography variant="h6" gutterBottom>
        Филиалы
      </Typography>
      <StyledForm>
        <Typography variant="subtitle1">Добавить новый филиал</Typography>
        <RadioGroup
          row
          value={shopType}
          onChange={(e) => setShopType(e.target.value)}
        >
          <FormControlLabel
            value="online"
            control={<Radio />}
            label="Онлайн-магазин"
          />
          <FormControlLabel
            value="offline"
            control={<Radio />}
            label="Офлайн-магазин"
          />
          <FormControlLabel
            value="marketplace"
            control={<Radio />}
            label="Магазин на маркетплейсе"
          />
        </RadioGroup>
        <TextField
          label="Название компании"
          fullWidth
          value={newShopName}
          onChange={(e) => setNewShopName(e.target.value)}
        />
        <TextField
          label="Описание компании"
          fullWidth
          value={newShopDescription}
          onChange={(e) => setNewShopDescription(e.target.value)}
          multiline
          rows={3}
        />
        <TextField
          label="Ссылка на страницу магазина"
          fullWidth
          value={newShopUrl}
          onChange={(e) => setNewShopUrl(e.target.value)}
        />
        <Button
          variant="contained"
          color="primary"
          onClick={createNewShop}
          sx={{ alignSelf: "flex-start" }}
        >
          Добавить филиал
        </Button>
      </StyledForm>
      <Divider sx={{ marginY: 3 }} />
      <Grid container spacing={2}>
        {shops.map((shop) => (
          <Grid item xs={12} sm={6} md={4} key={shop.getShopId()}>
            <Link
              to={`/shop/${shop.getShopId()}`} // Добавляем ссылку на магазин
              style={{ textDecoration: "none" }}
            >
              <StyledCard>
                <Typography variant="subtitle1" gutterBottom>
                  {shop.getName()}
                </Typography>
                <Chip
                  label={
                    "online"
                    // shop.getType() === "online"
                    //   ? "Онлайн-магазин"
                    //   : shop.getType() === "offline"
                    //   ? "Офлайн-магазин"
                    //   : "Магазин на маркетплейсе"
                  }
                  color="primary"
                  variant="outlined"
                  sx={{ marginBottom: 1 }}
                />
                <Typography variant="body2" color="textSecondary">
                  {shop.getDescription()}
                </Typography>
              </StyledCard>
            </Link>
          </Grid>
        ))}
      </Grid>
    </StyledPage>
  );
};

export default ShopPage;
