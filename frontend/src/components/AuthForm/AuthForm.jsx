import { useState } from 'react';
import { Button } from '@/components/Button';
import { Input } from '@/components/Input';
import * as SC from './AuthForm.styles';
import {useDispatch, useSelector} from "react-redux";
import {setIsLoggedIn, setIsLogin, setProfileEmail, setProfileName} from "@/store/appSlice.js";
import {getUserInfoRequest, loginRequest, registerRequest} from "@/api/api.js";
import {useNavigate} from "react-router-dom";
import {ROUTES} from "@/app/Router/routes.js";

export const AuthForm = () => {
    const isLogin = useSelector(state => state.isLogin);
    const dispatch = useDispatch();
    const navigate = useNavigate();
    const [formData, setFormData] = useState({
        email: '',
        username: '',
        password: '',
        confirmPassword: ''
    });

    const setToken = (token) => {
        localStorage.setItem('accessToken', token.accessToken);
        localStorage.setItem('refreshToken', token.refreshToken);

        dispatch(setIsLoggedIn(true));
        dispatch(setProfileName(formData.username));

        navigate(ROUTES.root);
    };

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData(prev => ({ ...prev, [name]: value }));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();

        let errorMessage = '';

        if (formData.email.length < 5 || formData.email.length > 255) {
            errorMessage += 'Длина почты должна быть от 5 до 255.\n';
        }
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if (!emailRegex.test(formData.email)) {
            errorMessage += 'Неверный формат почты.\n';
        }

        if (formData.password.length < 8 || formData.password.length > 64) {
            errorMessage += 'Длина пароля должна быть от 8 до 64.\n';
        }
        const passwordRegex = /^[A-Za-z0-9#?!@$%^&*-]+$/;
        if (!passwordRegex.test(formData.password)) {
            errorMessage += 'Неверный формат пароля. Пароль может содержать латинские буквы, цифры и специальные символы #?!@$%^&*-\n';
        }

        if(errorMessage.length > 0) {
            alert(errorMessage);
            return;
        }

        if(isLogin) {
            const token = await loginRequest(formData.email, formData.password);
            if (token) {
                setToken(token);
                const userInfo = await getUserInfoRequest();
                dispatch(setProfileName(userInfo.username));
                dispatch(setProfileEmail(userInfo.email));
            }
            else {
                alert('Ошибка при входе.');
            }
            return;
        }

        if (formData.username.length < 3 || formData.username.length > 50) {
            errorMessage += 'Длина логина должна быть от 3 до 50.\n';
        }
        const usernameRegex = /^[a-zA-Z0-9_]+$/;
        if (!usernameRegex.test(formData.username)) {
            errorMessage += 'Неверный формат логина. Логин может содержать латинские буквы, цифры и нижние подчёркивания.\n';
        }

        if (formData.password !== formData.confirmPassword) {
            errorMessage += 'Пароли не совпадают.\n';
        }

        if(errorMessage.length > 0) {
            alert(errorMessage);
            return;
        }

        const token = await registerRequest(formData.email, formData.username, formData.password);
        if (token) {
            setToken(token);
        }
        else {
            alert('Ошибка при регистрации.');
        }
    };

    return (
        <SC.AuthContainer>
            <SC.AuthForm onSubmit={handleSubmit}>
                <SC.AuthFormTitle>{isLogin ? 'Вход' : 'Регистрация'}</SC.AuthFormTitle>
                <SC.AuthFormSubtitle>Добро пожаловать в Article-Alchemy</SC.AuthFormSubtitle>

                {!isLogin && (
                    <Input
                        label="Email"
                        name="email"
                        value={formData.email}
                        onChange={handleChange}
                        required
                    />
                )}

                <Input
                    label={isLogin ? 'Email' : 'Логин'}
                    name={isLogin ? 'email' : 'username'}
                    value={isLogin ? formData.email : formData.username}
                    onChange={handleChange}
                    required
                />

                <Input
                    label="Пароль"
                    name="password"
                    type="password"
                    value={formData.password}
                    onChange={handleChange}
                    required
                />

                {!isLogin && (
                    <Input
                        label="Повторите пароль"
                        name="confirmPassword"
                        type="password"
                        value={formData.confirmPassword}
                        onChange={handleChange}
                        required
                    />
                )}

                <Button content={isLogin ? 'Войти' : 'Зарегистрироваться'} variant={'authButton'}/>

                <SC.ToggleText>
                    {isLogin ? 'Все еще нет аккаунта?' : 'Уже есть аккаунт?'}
                    <SC.ToggleButton
                        onClick={() => dispatch(setIsLogin(!isLogin))}
                    >
                        {isLogin ? 'Зарегистрироваться' : 'Войти'}
                    </SC.ToggleButton>
                </SC.ToggleText>
            </SC.AuthForm>
        </SC.AuthContainer>
    );
};