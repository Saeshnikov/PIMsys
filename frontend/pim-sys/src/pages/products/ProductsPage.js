import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import {
  Typography,
  Box,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Chip,
  Button,
} from "@mui/material";
import { ShopClient } from "../../grpc/shop/shop_grpc_web_pb";
import { ListProductsRequest } from "../../grpc/shop/shop_pb";

const ShopProductsPage = () => {
  const { shopId } = useParams();
  const [products, setProducts] = useState([]);
  const [shopName, setShopName] = useState("");

//   const shopClient = new ShopClient("http://localhost:8001", null, null);

//   useEffect(() => {
//     const token = localStorage.getItem("jwt_token");
//     const metadata = { authorization: token };

//     // Загружаем товары магазина
//     const request = new ListProductsRequest();
//     request.setShopId(shopId);

//     shopClient.listProducts(request, metadata, (err, response) => {
//       if (err) {
//         console.error("Error loading products", err);
//       } else {
//         setProducts(response.getProductsList());
//         setShopName(response.getShopName());
//       }
//     });
//   }, [shopId]);

  return (
    <Box sx={{ padding: 4 }}>
      <Typography variant="h4" gutterBottom>
        Товары магазина: {shopName}
      </Typography>
      <Button
        variant="outlined"
        onClick={() => window.history.back()}
        sx={{ marginBottom: 2 }}
      >
        Назад
      </Button>
      <TableContainer>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Наличие</TableCell>
              <TableCell>Категория</TableCell>
              <TableCell>Название</TableCell>
              <TableCell>Цена</TableCell>
              <TableCell>Количество</TableCell>
              <TableCell>Описание</TableCell>
              <TableCell>Статус</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {products.map((product) => (
              <TableRow key={product.getId()}>
                <TableCell>
                  {/* {product.getAvailable() ? "Есть" : "Нет"} */}
                </TableCell>
                <TableCell>{product.getCategory()}</TableCell>
                <TableCell>{product.getName()}</TableCell>
                <TableCell>{product.getPrice()} ₽</TableCell>
                <TableCell>{product.getQuantity()}</TableCell>
                <TableCell>{product.getDescription()}</TableCell>
                <TableCell>
                  <Chip
                    label={product.getStatus()}
                    color={
                        "В продаже"
                    //   product.getStatus() === "В продаже"
                    //     ? "success"
                    //     : product.getStatus() === "Закончился"
                    //     ? "error"
                    //     : "default"
                    }
                  />
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  );
};

export default ShopProductsPage;
