import * as SC from './SearchInput.styles';
import { Button } from '@/components/Button';
import {decrementFreeRequestsCount, setAnswerContent} from "@/store/appSlice.js";
import {useDispatch, useSelector} from "react-redux";
import {useState} from "react";
import {getHistory, registeredUserSummaryRequest, unregisteredUserSummaryRequest} from "@/api/api.js";

export const SearchInput = ( { backgroundColorButton='var(--color-grape)', colorButton='white', contentButton='Вперед!' } ) => {
    const isLoggedIn = useSelector((state) => state.isLoggedIn);
    const freeRequestsCount = useSelector((state) => state.freeRequestsCount);
    const dispatch = useDispatch();


    const [link, setLink] = useState('');
    const [email, setEmail] = useState('');

    const handleSubmit = async () => {

        let errorMessage = '';
        if (email.length < 5 || email.length > 255) {
            errorMessage += 'Длина почты должна быть от 5 до 255.\n';
        }
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if (!emailRegex.test(email)) {
            errorMessage += 'Неверный формат почты.\n';
        }
        if(link.length < 1) {
            errorMessage += 'Нет ссылки!';
        }

        if(errorMessage.length !== 0 && !isLoggedIn) {
            alert(errorMessage);
            return;
        }

        if(freeRequestsCount === 0 && !isLoggedIn) {
            dispatch(setAnswerContent({data: null, status: 'none'}));
            return;
        }

        if(isLoggedIn) {
            dispatch(setAnswerContent({data: 'Ожидание ответа...', status: 'pending'}));
            const result = await registeredUserSummaryRequest(link);
            if(result) {
                let counter = 0;
                let history = await getHistory();
                let currentAnswer = history.items.find((item) => item.contentId == result.contentId);
                while(currentAnswer.data.length === 0) {
                    await new Promise(resolve => setTimeout(resolve, 100));
                    history = await getHistory();
                    currentAnswer = history.items.find((item) => item.contentId == result.contentId);
                    counter += 1;
                    if(counter === 50) {
                        break;
                    }
                }

                if(currentAnswer.data.length > 0) {
                    dispatch(setAnswerContent({body: JSON.parse(currentAnswer.data), status: 'completed'}));
                    return;
                }
                dispatch(setAnswerContent({body: null, status: 'pending'}));
                return;
            }
            dispatch(setAnswerContent({body: null, status: 'failed'}));
        }
        else {
            dispatch(decrementFreeRequestsCount());
            if(freeRequestsCount !== 0) {
                dispatch(setAnswerContent(
                    {
                        body: {
                            main_title: 'Ответ придет на указанный email.',
                            data: []

                        },
                        status: 'completed'
                    }));

                const result = await unregisteredUserSummaryRequest(link, email);
                if(!result) {
                    dispatch(setAnswerContent({body: null, status: 'failed'}));
                }
            }
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