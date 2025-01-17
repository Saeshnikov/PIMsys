import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import DeleteForeverIcon from '@mui/icons-material/DeleteForever';
import { IconButton } from '@mui/material';
import {
  Button,
  TextField,
  Box,
  Typography,
  Grid,
  Divider,
  RadioGroup,
  Radio,
  FormControlLabel,
  Card,
  Chip,
} from "@mui/material";
import { styled } from "@mui/material/styles";
import Checkbox from '@mui/material/Checkbox';
import { TemplateClient } from "../../grpc/template/template_grpc_web_pb";
import {
  ListTemplatesRequest,
  ListTemplatesResponse,
  NewTemplateRequest,
  NewTemplateResponse,
  DeleteTemplateRequest,
  DeleteTemplateResponse,
  TemplateInfo,
  AttributeInfo,
} from "../../grpc/template/template_pb";
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

const CategoryPage = () => {
  const { branchId } = useParams();
  const navigate = useNavigate();
  const [templates, setTemplates] = useState([]);
  const [newTemplate, setNewTemplate] = useState({name:"",description:"", branchId:""});
  const [newAttributes, setNewAttributes] = useState([]);
  const deleteIcon = <DeleteForeverIcon />;

  const templateClient = new TemplateClient("http://localhost:8004", null, null);

  const loadTemplates = () => {
    const request = new ListTemplatesRequest();
    request.setBranchId(branchId);
    const token = localStorage.getItem("jwt_token");
    const metadata = {
      authorization: token,
    };
    templateClient.listTemplates(request, metadata, (err, response) => {
      if (err) {
        console.error("Error loading branches", err);
      } else {
        setTemplates(response.getInfoList().map((template) => template.toObject()));
      }
    });
  };


  const createNewTemplate = () => {
    const request = new NewTemplateRequest();
    request.setBranchId(branchId)
    request.setName(newTemplate.name);
    request.setDescription(newTemplate.description)
    var tmpAttributes = new Array();
    newAttributes.map((attr) => {
      const newAttr = new AttributeInfo();
      newAttr.setType(attr.type);
      newAttr.setIsValueRequired(attr.is_value_required);
      newAttr.setIsUnique(attr.is_unique);
      newAttr.setName(attr.name);
      newAttr.setDescription(attr.description);
      tmpAttributes.push(newAttr);
    });
    request.setAttributesList(tmpAttributes);
    const token = localStorage.getItem("jwt_token");
    const metadata = {
      authorization: token,
    };
    
    templateClient.newTemplate(request, metadata, (err, response) => {
      if (err) {
        console.error("Error creating template", err);
      } else {
        loadTemplates();
        setNewTemplate({name:"",description:"", branchId:""});
        setNewAttributes([]);
      }
    });
  };

  const deleteTemplate = (templateId) => {
    const request = new DeleteTemplateRequest();
    request.setTemplateId(templateId);

    const token = localStorage.getItem("jwt_token");
    const metadata = {
      authorization: token,
    };

    templateClient.deleteTemplate(request, metadata, (err, response) => {
      if (err) {
        console.error("Error deleting филиала", err);
      } else {
        loadTemplates();
      }
    });
  };

  useEffect(() => {
    const token = localStorage.getItem("jwt_token");

    if (!token) {
      navigate("/");
      return;
    }
    loadTemplates();
  }, []);

  return (
    <div>
    <StyledPage>
      <Typography variant="h4" gutterBottom>
        Категории
      </Typography>
      <StyledForm>
        <Typography variant="subtitle1">Добавить новую категорию</Typography>
        <TextField
          label="Название категории"
          fullWidth
          value={newTemplate.name}
          onChange={(e) => setNewTemplate({ ...newTemplate, name: e.target.value })}
        />
        <TextField
          label="Описание категории"
          fullWidth
          value={newTemplate.description}
          onChange={(e) => setNewTemplate({ ...newTemplate, description: e.target.value })}
          multiline
          rows={3}
        />
        {newAttributes.map((attribute) => (
          <div>
            <Divider variant="middle" sx={{ marginY: 1 }}/>
            <TextField
              label="Название атрибута"
              fullWidth
              value={attribute.name}
              onChange={(e) => setNewAttributes((prevAttributes) =>
                prevAttributes.map((attr) =>
                  attr.id === attribute.id
                    ? { ...attr, name: e.target.value } // Создаём новый объект
                    : attr
                )
              )}
            />
            <TextField
              label="Описание атрибута"
              fullWidth
              value={attribute.description}
              onChange={(e) => setNewAttributes((prevAttributes) =>
                prevAttributes.map((attr) =>
                  attr.id === attribute.id
                    ? { ...attr, description: e.target.value } // Создаём новый объект
                    : attr
                )
              )}
            />
            <RadioGroup
              row
              value={attribute.type}
              onChange={(e) => setNewAttributes((prevAttributes) =>
                prevAttributes.map((attr) =>
                  attr.id === attribute.id
                    ? { ...attr, type: e.target.value } // Создаём новый объект
                    : attr
                )
              )}
            >
              <FormControlLabel
                value="text"
                control={<Radio />}
                label="Текст"
              />
              <FormControlLabel
                value="number"
                control={<Radio />}
                label="Число"
              />
              <FormControlLabel
                value="boolean"
                control={<Radio />}
                label="бул"
              />
            </RadioGroup>
            <Typography variant="subtitle2" gutterBottom>
              Значение обязательно
            </Typography>
            <Checkbox
              label="Значение обязательно"
              value= {attribute.is_value_required}
              onChange={(e) => setNewAttributes((prevAttributes) =>
                prevAttributes.map((attr) =>
                  attr.id === attribute.id
                    ? { ...attr, is_value_required: e.target.value } // Создаём новый объект
                    : attr
                )
              )}
            />
            <Typography variant="subtitle2" gutterBottom>
              Значение уникально
            </Typography>
            <Checkbox
              label="Значение уникально"
              value= {attribute.is_unique}
              onChange={(e) => setNewAttributes((prevAttributes) =>
                prevAttributes.map((attr) =>
                  attr.id === attribute.id
                    ? { ...attr, is_unique: e.target.value } // Создаём новый объект
                    : attr
                )
              )}
            />
          </div>
        ))}
        <Button
          variant="contained"
          color="primary"
          onClick={(e) => setNewAttributes((prevAttributes) => [
            ...prevAttributes, // Создаём новый массив
            {
              id: prevAttributes.length,
              type: "text",
              is_value_required: false,
              is_unique: false,
              name: "",
              description: "",
            },
          ])}
          sx={{ alignSelf: "flex-start" }}
        >
          + Добавить аттрибут
        </Button>
        <Button
          variant="contained"
          color="primary"
          onClick={createNewTemplate}
          sx={{ alignSelf: "flex-start" }}
        >
          Добавить категорию
        </Button>
      </StyledForm>
      <Divider sx={{ marginY: 3 }} />
      <Grid container spacing={2}>
        {templates.map((template) => (
          <Grid item xs={12}  key={template.templateId}>
            <StyledCard>
              <Grid container spacing={12}>
                <Grid item xs={7}>
                    <Typography variant="subtitle1" gutterBottom>
                      {template.name}
                    
                    </Typography>
                  </Grid>
                  <Grid item xs={5}>
                    <IconButton onClick={() => deleteTemplate(template.templateId)}
                    variant="filled"
                    color="danger"
                    sx={{ alignSelf: "flex-end" }}
                    >{deleteIcon}</IconButton>
                  </Grid>
                </Grid>
                <Typography variant="body2" color="textSecondary">
                  {template.description}
                </Typography> 
                {template.attributesList.map((attribute) => (
                  <div>
                    <Divider variant="middle" sx={{ marginY: 1 }}/>
                    <Typography variant="body2" color="textSecondary">
                      {attribute.name}
                    </Typography> 
                    <Typography variant="body2" color="textSecondary">
                      {attribute.description}
                    </Typography> 
                    <Chip
                      label={attribute.type}
                      color="success"
                      variant="outlined"
                      sx={{ marginBottom: 1 }}
                    />
                    <div></div>
                    {attribute.isValueRequired &&
                      <Chip
                        label="required"
                        color="primary"
                        variant="outlined"
                        sx={{ marginBottom: 1 }}
                      />
                    }
                    {attribute.isUnique &&
                      <Chip
                        label="unique"
                        color="primary"
                        variant="outlined"
                        sx={{ marginBottom: 1 }}
                      />
                    }
                  </div>
                ))}
            </StyledCard>
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

export default CategoryPage;
