import React, { useState, useEffect } from "react";
import { ProductClient } from "../../grpc/products/products_grpc_web_pb"; // Клиент gRPC
import { TemplateClient } from "../../grpc/template/template_grpc_web_pb"; // Клиент gRPC
import { useParams } from "react-router-dom";
import { 
  ProductInfo, 
  Attribute, 
  ProductInfoWithId, 
  DeleteProductRequest, 
  Products,
  Empty,
} from "../../grpc/products/products_pb"; // Сгенерированные сообщения
import { 
  ListTemplatesResponse,
  ListTemplatesRequest,
  TemplateInfo,
  AttributeInfo,
} from "../../grpc/template/template_pb"; // Сгенерированные сообщения
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
  const templateClient = new TemplateClient("http://localhost:8004"); // URL gRPC-сервера
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
        console.error("Ошибка загрузки списка продуктов:", err.message);
        return;
      }
      setProducts(response.getProductList().map((product) => product.toObject()));
      
    });

    // console.log("newCategory")
    // console.log(newCategory)
    // console.log("products")
    // console.log(products)
    // console.log("rows")
    // console.log(rows)
    
    const templateRequest = new ListTemplatesRequest();
    templateRequest.setBranchId(branchId);
    templateClient.listTemplates(templateRequest, metadata, (err, response) => {
      if (err) {
        console.error("Ошибка загрузки списка категорий:", err.message);
        return;
      }
      setCategories(response.getInfoList().map((template) => {
        const tmp = template.toObject();
        // console.log(tmp)
        // console.log(tmp.product.attributesList)
        return {categoryId: tmp.templateId,
          name: tmp.name,
          description: tmp.description,
          attributes: tmp.attributesList}
      }));
    })
  };

  useEffect(() => {
    handleSetRows(newCategory);
  }, [newCategory, products]);


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

  const handleSetRows = (cat) => {
    if(!cat){
      return
    }
    setRows(products.filter(
      (p)=> {return p.product.categoryId === cat.categoryId}
    ).map((p) => {
        return {id: p.productId,
        status: p.product.status,
        name: p.product.name,
        price: p.product.price,
        amount: p.product.amount,
        branchId: p.product.branchId,
        categoryId: p.product.categoryId,
        attributes: p.product.attributesList}
      // console.log(tmp)
      // console.log(tmp.product.attributesList)
      
    }))
  }

  const handleDeleteProduct = async (Product_id) => {
    const request = new DeleteProductRequest();
    // console.log(Product_id)
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
            if (!filteredAttr){ return ""}
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
              onChange={(e) => {setNewProduct({ ...newProduct, name: e.target.value });}}
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
