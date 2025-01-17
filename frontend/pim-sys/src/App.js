import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import RegisterPage from "./pages/register/RegisterPage"; // Компонент регистрации
import SuccessPage from "./pages/SuccessPage"; // Страница успеха
// import { AuthClient } from "./grpc/sso/sso_grpc_web_pb";
import "./App.css";
import ShopPage from "./pages/shop/ShopPage";
import LogsPage from "./pages/logs/LogsPage";
import BranchPage from "./pages/branch/BranchPage";
import ProductsPage from "./pages/products/ProductsPage";
import CategoryPage from "./pages/category/CategoryPage";

// const authClient = new AuthClient("http://localhost:8000");

function App() {
  return (
    <Router>
      <Routes>
        <Route
          path="/"
          element={<RegisterPage/>}
        />
        {/* <Route path="/success" element={<SuccessPage />} /> */}
        <Route path="/shop" element={<ShopPage />} />
        <Route path="/logs" element={<LogsPage />} />
        <Route path="/shop/:shopId/branches" element={<BranchPage />} />
        <Route path="/shop/:shopId/:branchId/products" element={<ProductsPage />} />
        <Route path="/shop/:shopId/:branchId/category" element={<CategoryPage />} />
      </Routes>
    </Router>
  );
}

export default App;
