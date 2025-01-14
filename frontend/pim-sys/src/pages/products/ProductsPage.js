import React, { useState, useEffect } from "react";
import { ProductClient } from "../../grpc/products/products_grpc_web_pb"; // Клиент gRPC
import { useParams } from "react-router-dom";
import { 
  ProductInfo, 
  Attribute, 
  ProductInfoWithId, 
  DeleteProductRequest, 
  Products,
  Empty,
} from "../../grpc/products/products_pb"; // Сгенерированные сообщения
import { useNavigate } from "react-router-dom";
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import DeleteForeverIcon from '@mui/icons-material/DeleteForever';
import EditIcon from '@mui/icons-material/Edit';
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

const ProductsPage = () => {
  const client = new ProductClient("http://localhost:8002"); // URL gRPC-сервера
  const { branchId } = useParams();

  const [products, setProducts] = useState([]);
  const [newAttributes, setNewAttributes] = useState([]);
  const [categories, setCategories] = useState([]);
  const [newProduct, setNewProduct] = useState({ name: "", price: 0, amount: 0, status: "stock",attributes:newAttributes});
  const [editProduct, setEditProduct] = useState(null);
  const [newCategory, setNewCategory] = useState(null);
  const [rows, setRows] = useState([]);
  const [columns, setColumns] = useState([]);

  const deleteIcon = <DeleteForeverIcon />;
  const editIcon = <EditIcon />;
  
  useEffect(() => {
    fetchProducts();
  }, []);

  const navigate = useNavigate(); // Хук для управления навигацией

  const token = localStorage.getItem("jwt_token");
  const metadata = {
    authorization: token,
  };

  const fetchProducts = async () => {
    const request = new Empty();
    client.listProducts(request, metadata, (err, response) => {
      if (err) {
        console.error("Ошибка загрузки списка магазинов:", err.message);
        return;
      }
      setProducts(response.getProductList().map((product) => product.toObject()));
      setRows(response.getProductList().map((product) => {
        const tmp = product.toObject();
        // console.log(tmp)
        // console.log(tmp.product.attributesList)
        return {id: tmp.productId,
          status: tmp.product.status,
          name: tmp.product.name,
          price: tmp.product.price,
          amount: tmp.product.amount,
          branchId: tmp.product.branchId,
          categoryId: tmp.product.categoryId,
          attributes: tmp.product.attributesList}
      }))
    });
    setCategories([{ categoryId:1,name: "test-category", description: "test-description",attributes:[{id:1,type:"text",name: "test-attribute-text"},{id:2,type:"number",name: "test-attribute-numeric"}]}])
  };

  const handleAddProduct = async () => {
    const request = new ProductInfo();
    request.setName(newProduct.name);
    request.setStatus(newProduct.status);
    request.setAmount(newProduct.amount);
    request.setPrice(newProduct.price);
    request.setBranchId(branchId);
    request.setCategoryId(newCategory.categoryId);
    request.setAttributesList(newAttributes);

    client.newProduct(request, metadata, (err, response) => {
      if (err) {
        console.error("Ошибка добавления продукта:", err.message);
        return;
      }
      // setNewAttributes(null);
      setNewProduct({ name: "", price: 0, amount: 0, status: "stock",attributes:newAttributes});
      fetchProducts();
    });
  };

  const handleSetEditProduct = async(params) => {
    const forEdit = new ProductInfoWithId();
    const productInfo = new ProductInfo();
    productInfo.setName(params.name);
    productInfo.setStatus(params.status);
    productInfo.setAmount(params.amount);
    productInfo.setPrice(params.price);
    productInfo.setBranchId(branchId);
    productInfo.setCategoryId(params.categoryId);
    forEdit.setProduct(productInfo)
    forEdit.setProductId(params.id)

    setEditProduct(forEdit)
  }

  const handleEditProduct = async () => {
    const request = new ProductInfoWithId();
    const productInfo = new ProductInfo();
    productInfo.setName(editProduct.name);
    productInfo.setStatus(editProduct.status);
    productInfo.setAmount(editProduct.amount);
    productInfo.setPrice(editProduct.price);
    productInfo.setBranchId(branchId);
    productInfo.setCategoryId(editProduct.categoryId);
    request.setProduct(productInfo)
    request.setProductId(editProduct.id)
    

    client.alterProduct(request, metadata, (err, response) => {
      if (err) {
        console.error("Ошибка изменения продукта:", err.message);
        alert("Ошибка изменения продукта:", err.message)
        return;
      }
      setEditProduct(null);
      fetchProducts();
    });
  };

  const handleDeleteProduct = async (Product_id) => {
    const request = new DeleteProductRequest();
    console.log(Product_id)
    request.setProductId(Product_id);

    client.deleteProduct(request, metadata, (err, response) => {
      if (err) {
        console.error("Ошибка удаления продукта:", err.message);
        alert("Ошибка удаления продукта:", err.message)
        return;
      }
      fetchProducts();
    });
  };

  const defaultColumns = [
    { field: 'name', headerName: 'Название', width: 120 },
    { field: 'price', headerName: 'Цена', width: 100 },
    { field: 'amount', headerName: 'Количество', width: 120 },
    {
      field: 'status',
      headerName: 'Статус',
      width: 150,
      renderCell: (params) => (
        <Chip
          label={
            params.value === "stock"
              ? "В продаже"
              : params.value === "archive"
              ? "В архиве"
              : params.value === "withdrawn_from_sale"
              ? "Снято с продажи"
              : "Нет в наличии"
          }
          color="primary"
          variant="outlined"
          sx={{ marginBottom: 1 }}
        />
      ),
    }];

  const rowsButtons = {
    field: 'actions',
    headerName: '',
    width: 200,
    renderCell: (params) => (
      <div style={{ display: 'flex', gap: '10px' }}>
        <IconButton onClick={() => setEditProduct(params.row)}
          variant="filled"
          color="primary"
          sx={{ alignSelf: "flex-end" }}
        >{editIcon}</IconButton>
        <IconButton onClick={() => handleDeleteProduct(params.row.id)}
          variant="filled"
          color="danger"
          sx={{ alignSelf: "flex-end" }}
        >{deleteIcon}</IconButton>
      </div>
    ),
  };

  const handleNewCategory = async (e) => {
    setNewCategory(e);
    var tmpAttributes = new Array();
    var tmpColumns = new Array();
    tmpColumns.push(...defaultColumns);
    e.attributes?.map((attribute) => {
      const attr = new Attribute();
      attr.setId(attribute.id);
      tmpAttributes.push(attr);
      tmpColumns.push(
        { field: `${attribute.id}`,
          headerName: attribute.name,
          valueGetter:  (value, row) => {
            const filteredAttr = row.attributes
            ?.find((attr) => attr.id == attribute.id);
            return attribute.type === "text" ?
              `${filteredAttr.valueText.toString()}`
              : attribute.type === "number"
              ? `${filteredAttr.valueNumber.toString()}`
              : `${filteredAttr.valueBool.toString()}`;
          }
      },

    );
    });
    tmpColumns.push(rowsButtons);

    setNewAttributes(tmpAttributes);
    setColumns(tmpColumns);
  };

  const handleNewAttributes = async (attributeId,value) => {
    var tmpAttribute = newAttributes
    
    tmpAttribute.map((attribute) => {
      if (attribute.getId() === attributeId){
        newCategory.attributes?.map((attr) => {
          if (attr.id === attributeId){
            attr.type === "text"
            ? attribute.setValueText(value)
            : attr.type === "number"
            ? attribute.setValueNumber(value)
            : attribute.setValueBool(value)
            
          }
        })
      }
    })
    
    setNewAttributes(tmpAttribute)
  };

  // const handleOpenBranches = async (Product) => {
  //   const request = new ProductInfo();
  //   localStorage.setItem("Product_name", Product.name);
  //   localStorage.setItem("Product_description", Product.description);

  //   navigate(`/Product/${Product.ProductId}/branches`);
  // };

  

  return (
    <div >
        <StyledPage sx={{width: '100%'}}>
          <Divider sx={{ marginY: 3 }}/>
          <Typography variant="h6" gutterBottom>
            Продукты
          </Typography>
          <FormControl fullWidth>
            <InputLabel>Category</InputLabel>
            <Select
              value={newCategory}
              label="Category"
              onChange={(e) => handleNewCategory(e.target.value)}
              renderValue={(selected) => (
                selected.name
              )}
            >
              {categories.map((category) => (
                <MenuItem value={category}>{category.name}</MenuItem>
              ))}
              
            </Select>
          </FormControl>
          {newCategory != null && <StyledForm>
            <Typography variant="subtitle1">Добавить новый продукт</Typography>
            <TextField
              label="Название продукта"
              fullWidth
              value={newProduct.name}
              onChange={(e) => {setNewProduct({ ...newProduct, name: e.target.value });console.log(newCategory.name) ;}}
            />
            <TextField
              label="Цена"
              fullWidth
              value={newProduct.price}
              onChange={(e) => setNewProduct({ ...newProduct, price: e.target.value })}
            />
            <TextField
              label="Количество"
              fullWidth
              value={newProduct.amount}
              onChange={(e) => setNewProduct({ ...newProduct, amount: e.target.value })}
            />
            <RadioGroup
              row
              value={newProduct.status}
              onChange={(e) => setNewProduct({ ...newProduct, status: e.target.value })}
              >
                <FormControlLabel
                  value="stock"
                  control={<Radio />}
                  label="В продаже"
                />
                <FormControlLabel
                  value="archive"
                  control={<Radio />}
                  label="В архиве"
                />
                <FormControlLabel
                  value="withdrawn_from_sale"
                  control={<Radio />}
                  label="Снято с продажи"
                />
                <FormControlLabel
                  value="out_of_stock"
                  control={<Radio />}
                  label="Нет в наличии"
                />
            </RadioGroup>
             {newCategory.attributes?.map((attribute) => (
              <TextField
              label={attribute.name}
                fullWidth
                onChange={(e) => handleNewAttributes(attribute.id, e.target.value )}
              />
            ))}
            <Button
              variant="contained"
              color="primary"
              onClick={handleAddProduct}
              sx={{ alignSelf: "flex-start" }}
            >
              Добавить продукт
            </Button>
          </StyledForm>}



          {/* <Divider sx={{ marginY: 3 }} />
          <Grid container spacing={2}>
            {products.map((product) => (
              <Grid item xs={12} sm={6} md={4} key={product.productId}>
                  {(editProduct==null || editProduct.productId!==product.productId) && (<StyledCard onClick={() => setEditProduct(product)}>
                    <Typography variant="subtitle1" gutterBottom>
                      {product.product.name}
                    </Typography>
                    <Chip
                      label={
                        product.product.status === "stock"
                          ? "В продаже"
                          : product.product.status === "archive"
                          ? "В архиве"
                          : product.product.status === "withdrawn_from_sale"
                          ? "Снято с продажи"
                          : "Нет в наличии"
                      }
                      color="primary"
                      variant="outlined"
                      sx={{ marginBottom: 1 }}
                    />
                    <Typography variant="body2" color="textSecondary">
                      Количество: {product.product.amount}
                    </Typography>
                    <Typography variant="body2" color="textSecondary">
                      Цена: {product.product.price}
                    </Typography>
                  </StyledCard>
                  )}
                  {editProduct!=null && editProduct.productId===product.productId && (
                  <div>
                    <StyledCard>
                      <TextField
                        label="Название продукта"
                        fullWidth
                        value={newProduct.name}
                        onChange={(e) => setNewProduct({ ...newProduct, name: e.target.value })}
                      />
                      <TextField
                        label="Цена"
                        fullWidth
                        value={newProduct.price}
                        onChange={(e) => setNewProduct({ ...newProduct, price: e.target.value })}
                        multiline
                        rows={3}
                      />
                      <TextField
                        label="Количество"
                        fullWidth
                        value={newProduct.amount}
                        onChange={(e) => setNewProduct({ ...newProduct, amount: e.target.value })}
                      />
                      <RadioGroup
                        row
                        value={newProduct.status}
                        onChange={(e) => setNewProduct({ ...newProduct, status: e.target.value })}
                        >
                          <FormControlLabel
                            value="stock"
                            control={<Radio />}
                            label="В продаже"
                          />
                          <FormControlLabel
                            value="archive"
                            control={<Radio />}
                            label="В архиве"
                          />
                          <FormControlLabel
                            value="withdrawn_from_sale"
                            control={<Radio />}
                            label="Снято с продажи"
                          />
                          <FormControlLabel
                            value="out_of_stock"
                            control={<Radio />}
                            label="Нет в наличии"
                          />
                      </RadioGroup>
                      
                      <Button onClick={handleEditProduct} 
                      variant="contained"
                      color="primary"
                      sx={{ alignSelf: "flex-start" }}
                      >Сохранить</Button>
                      <Button onClick={() => setEditProduct(null)}
                      variant="contained"
                      color="primary"
                      sx={{ alignSelf: "flex-start" }
                      }>Отмена</Button>
                    </StyledCard>
                    
                    
                  </div>
                )}
              </Grid>
            ))}
          </Grid> */}
          {editProduct!=null && newCategory != null && <StyledForm>
            <Typography variant="subtitle1">Изменить продукт</Typography>
            <TextField
              label="Название продукта"
              fullWidth
              value={editProduct.name}
              onChange={(e) => {setEditProduct({ ...editProduct, name: e.target.value });}}
            />
            <TextField
              label="Цена"
              fullWidth
              value={editProduct.price}
              onChange={(e) => setEditProduct({ ...editProduct, price: e.target.value })}
            />
            <TextField
              label="Количество"
              fullWidth
              value={editProduct.amount}
              onChange={(e) => setEditProduct({ ...editProduct, amount: e.target.value })}
            />
            <RadioGroup
              row
              value={editProduct.status}
              onChange={(e) => setEditProduct({ ...editProduct, status: e.target.value })}
              >
                <FormControlLabel
                  value="stock"
                  control={<Radio />}
                  label="В продаже"
                />
                <FormControlLabel
                  value="archive"
                  control={<Radio />}
                  label="В архиве"
                />
                <FormControlLabel
                  value="withdrawn_from_sale"
                  control={<Radio />}
                  label="Снято с продажи"
                />
                <FormControlLabel
                  value="out_of_stock"
                  control={<Radio />}
                  label="Нет в наличии"
                />
            </RadioGroup>
            <Button
              variant="contained"
              color="primary"
              onClick={handleEditProduct}
              sx={{ alignSelf: "flex-start" }}
            >
              Изменить
            </Button>
            <Button
              variant="contained"
              color="primary"
              onClick={() => setEditProduct(null)}
              sx={{ alignSelf: "flex-start" }}
            >
              Отмена
            </Button>
          </StyledForm>}
          <DataGrid rows={rows} columns={columns} pageSize={5} fullWidth/>
            
          </StyledPage>
          <Typography variant="body2" color="textSecondary">
              (c) Brigada, Inc
          </Typography>


    </div>

  );
};

export default ProductsPage;
