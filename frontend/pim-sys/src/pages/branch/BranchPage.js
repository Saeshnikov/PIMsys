import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { useParams } from "react-router-dom";
import LocalMallIcon from '@mui/icons-material/LocalMall';
import InterestsIcon from '@mui/icons-material/Interests';
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
import { BranchClient } from "../../grpc/branch/branch_grpc_web_pb";
import {
  ListBranchesRequest,
  NewBranchRequest,
  DeleteBranchRequest,
  AlterBranchRequest,
} from "../../grpc/branch/branch_pb";
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

const BranchPage = () => {
  const productIcon = <LocalMallIcon />;
  const categoriesIcon = <InterestsIcon />;
  const { shopId } = useParams();
  const navigate = useNavigate();
  const [shopName, setShopName] = useState("");
  const [shopDescription, setShopDescription] = useState("");
  const [branchId, setBranchId] = useState();
  const [newBranchName, setNewBranchName] = useState("");
  const [newBranchDescription, setNewBranchDescription] = useState("");
  const [newBranchAddress, setNewBranchAddress] = useState("");
  const [newBranchSite, setNewBranchSite] = useState("");
  const [branchType, setBranchType] = useState("online");
  const [branches, setBranches] = useState([]);

  const branchClient = new BranchClient("http://localhost:8003", null, null);

  const loadBranches = () => {
    // alert("Создано компанией Brigada INC");
    setShopName(localStorage.getItem("shop_name"));
    setShopDescription(localStorage.getItem("shop_description"));
    const request = new ListBranchesRequest();
    request.setShopId(shopId);
    const token = localStorage.getItem("jwt_token");
    const metadata = {
      authorization: token,
    };
    branchClient.listBranches(request, metadata, (err, response) => {
      if (err) {
        console.error("Error loading branches", err);
      } else {
        setBranches(response.getInfoList().map((branch) => branch.toObject()));
      }
    });
  };

  const createNewBranch = () => {
    const request = new NewBranchRequest();
    request.setShopId(shopId)
    request.setName(newBranchName);
    request.setDescription(newBranchDescription);
    request.setAddress(newBranchAddress);
    request.setSite(newBranchSite);
    request.setBranchType(branchType);
    const token = localStorage.getItem("jwt_token");
    const metadata = {
      authorization: token,
    };
    
    branchClient.newBranch(request, metadata, (err, response) => {
      if (err) {
        console.error("Error creating branch", err);
      } else {
        loadBranches();
        setNewBranchName("");
        setNewBranchDescription("");
        setNewBranchAddress("");
        setNewBranchSite("");
        setBranchType("online");
      }
    });
  };

  const deleteBranch = (branchId) => {
    const request = new DeleteBranchRequest();
    request.setBranchId(branchId);

    const token = localStorage.getItem("jwt_token");
    const metadata = {
      authorization: token,
    };

    branchClient.deleteBranch(request, metadata, (err, response) => {
      if (err) {
        console.error("Error deleting shop", err);
      } else {
        loadBranches();
      }
    });
  };

  useEffect(() => {
    const token = localStorage.getItem("jwt_token");

    if (!token) {
      navigate("/");
      return;
    }
    loadBranches();
  }, []);

  return (
    <div>
    <StyledPage>
      <Typography variant="h4" gutterBottom>
        Компания
      </Typography>
      <Typography variant="h5" gutterBottom>
        {shopName}
      </Typography>
      <Typography variant="body1" gutterBottom>
        {shopDescription}
      </Typography>
      <Divider sx={{ marginY: 3 }} />
      <Typography variant="h6" gutterBottom>
        Филиалы
      </Typography>
      <StyledForm>
        <Typography variant="subtitle1">Добавить новый филиал</Typography>
        <RadioGroup
          row
          value={branchType}
          onChange={(e) => setBranchType(e.target.value)}
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
          label="Название филиала"
          fullWidth
          value={newBranchName}
          onChange={(e) => setNewBranchName(e.target.value)}
        />
        <TextField
          label="Описание филиала"
          fullWidth
          value={newBranchDescription}
          onChange={(e) => setNewBranchDescription(e.target.value)}
          multiline
          rows={3}
        />
        <TextField
          label="Адрес филиала"
          fullWidth
          value={newBranchAddress}
          onChange={(e) => setNewBranchAddress(e.target.value)}
        />
        <TextField
          label="Сайт филиала"
          fullWidth
          value={newBranchSite}
          onChange={(e) => setNewBranchSite(e.target.value)}
        />
        <Button
          variant="contained"
          color="primary"
          onClick={createNewBranch}
          sx={{ alignSelf: "flex-start" }}
        >
          Добавить филиал
        </Button>
      </StyledForm>
      <Divider sx={{ marginY: 3 }} />
      <Grid container spacing={2}>
        {branches.map((branch) => (
          <Grid item xs={12} sm={6} md={4} key={branch.branchId}>
            {/* <Link
              to={`/branch/${branch.branchId}`} // Добавляем ссылку на магазин
              style={{ textDecoration: "none" }}
            > */}
              <StyledCard>
                <Typography variant="subtitle1" gutterBottom>
                  {branch.name}
                </Typography>
                
                <Chip
                  label={
                    branch.branchType === "online"
                      ? "Онлайн-магазин"
                      : branch.branchType === "offline"
                      ? "Офлайн-магазин"
                      : "Магазин на маркетплейсе"
                  }
                  color="primary"
                  variant="outlined"
                  sx={{ marginBottom: 1 }}
                />
                <div></div>
                <Chip
                  label="Продукты"
                  icon ={productIcon}
                  color="success"
                  variant="outlined"
                  sx={{ marginBottom: 1 }}
                />
                <Chip
                  label="Категории"
                  icon ={categoriesIcon}
                  color="success"
                  variant="outlined"
                  sx={{ marginBottom: 1 }}
                />
                <Typography variant="body2" color="textSecondary">
                  {branch.description}
                </Typography>
              </StyledCard>
            {/* </Link> */}
          </Grid>
        ))}
      </Grid>
      
      
    </StyledPage>
    <Typography variant="body2" color="textSecondary">
        (c) Brigada, Inc
    </Typography>
    </div>
  );
};

export default BranchPage;
