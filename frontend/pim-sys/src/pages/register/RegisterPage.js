import React from "react";
import RegisterForm from "./RegisterForm";
import LoginForm from "./LoginForm";
import { AuthClient } from "../../grpc/sso/sso_grpc_web_pb";
import "./App.css";
import Typography from "@mui/material/Typography";

const authClient = new AuthClient("http://localhost:8000");

const RegisterPage = () =>  {
  const [isLogin, setIsLogin] = React.useState(false); // Управление состоянием: регистрация или авторизация

  return (
    <div className="app-container">
      <div className="left-section">
        <h1>Добро пожаловать в PIM-Систему</h1>
        <ul>
          <li>
            <h3>Централизованное управление данными о товарах</h3>
            <p>
              Мы предоставляем единую платформу для хранения и обработки
              информации о ваших товарах. Это позволяет избежать ошибок и
              неточностей в данных.
            </p>
          </li>
          <li>
            <h3>Автоматизация процессов</h3>
            <p>
              Мы автоматизируем рутинные задачи, такие как обновление цен,
              изменение характеристик товаров и т.д. Это освобождает ваше время
              для более важных задач.
            </p>
          </li>
          <li>
            <h3>Анализ и отчётность</h3>
            <p>
              Мы предоставляем инструменты для анализа данных о продажах,
              остатках и других показателях. Это помогает вам принимать
              обоснованные решения.
            </p>
          </li>
        </ul>
      </div>
      <div className="right-section">
        {isLogin ? (
          <LoginForm
            authClient={authClient}
            onSwitchToRegister={() => setIsLogin(false)}
          />
        ) : (
          <RegisterForm
            authClient={authClient}
            onSwitchToLogin={() => setIsLogin(true)}
          />
        )}
      </div>
    </div>
  );
}

export default RegisterPage;
