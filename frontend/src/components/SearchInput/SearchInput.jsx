import * as SC from './SearchInput.styles';
import { Button } from '@/components/Button';
import {decrementFreeRequestsCount, setAnswerContent} from "@/store/appSlice.js";
import {useDispatch, useSelector} from "react-redux";
import {useState} from "react";

export const SearchInput = ( { backgroundColorButton='var(--color-grape)', colorButton='white', contentButton='Вперед!' } ) => {
    const isLoggedIn = useSelector((state) => state.isLoggedIn);
    const freeRequestsCount = useSelector((state) => state.freeRequestsCount);
    const dispatch = useDispatch();

    const [link, setLink] = useState('');
    const [email, setEmail] = useState('');

    const handleSubmit = () => {

        console.log('handleSubmit');
        let errorMessage = '';
        if (email.length < 5 || email.length > 255) {
            errorMessage += 'Длина почты должна быть от 5 до 255.\n';
        }
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if (!emailRegex.test(email)) {
            errorMessage += 'Неверный формат почты.\n';
        }

        if(errorMessage.length !== 0) {
            alert(errorMessage);
            return;
        }

        if(freeRequestsCount === 0 && !isLoggedIn) {
            dispatch(setAnswerContent('Нет ответа'));
            return;
        }

        console.log('запрос пересказа');
        if(isLoggedIn) {
            dispatch(setAnswerContent('типо ответ'));
        }
        if(!isLoggedIn) {
            dispatch(decrementFreeRequestsCount());
        }
    };

    const emailInput = isLoggedIn
        ?   <></>
        :   <SC.SearchInputContainer style={{marginTop: '5rem'}}>
                <SC.SearchInput placeholder={'Адрес почты...'}
                                value={email}
                                onChange={(e) => {
                                    setEmail(e.target.value);
                                }}
                                required />
            </SC.SearchInputContainer>;

    return (
        <>
            <SC.SearchInputContainer>
                <SC.SearchInput placeholder={'Ссылка на статью...'}
                                value={link}
                                onChange={(e) => {
                                    setLink(e.target.value);
                                }}
                                required />
                <Button backgroundColor={backgroundColorButton}
                        color={colorButton}
                        content={contentButton}
                        handleClick={handleSubmit}
                />
            </SC.SearchInputContainer>
            {emailInput}
        </>
    );
};