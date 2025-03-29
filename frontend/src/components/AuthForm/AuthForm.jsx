import { useState } from 'react';
import { Button } from '@/components/Button';
import { Input } from '@/components/Input';
import * as SC from './AuthForm.styles';

export const AuthForm = () => {
    const [isLogin, setIsLogin] = useState(true);
    const [formData, setFormData] = useState({
        email: '',
        username: '',
        password: '',
        confirmPassword: ''
    });

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData(prev => ({ ...prev, [name]: value }));
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        if (!isLogin && formData.password !== formData.confirmPassword) {
            alert('Пароли не совпадают!');
            return;
        }
        //...
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
                        onClick={() => setIsLogin(!isLogin)}
                    >
                        {isLogin ? 'Зарегистрироваться' : 'Войти'}
                    </SC.ToggleButton>
                </SC.ToggleText>
            </SC.AuthForm>
        </SC.AuthContainer>
    );
};