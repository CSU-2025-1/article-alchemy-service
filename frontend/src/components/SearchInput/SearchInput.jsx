import * as SC from './SearchInput.styles';
import { Button } from '@/components/Button';
import {decrementFreeRequestsCount, setAnswerContent} from "@/store/appSlice.js";
import {useDispatch, useSelector} from "react-redux";

export const SearchInput = ( { backgroundColorButton='var(--color-grape)', colorButton='white', contentButton='Вперед!' } ) => {
    const isLoggedIn = useSelector((state) => state.isLoggedIn);
    const freeRequestsCount = useSelector((state) => state.freeRequestsCount);
    const dispatch = useDispatch();
    return (
        <SC.SearchInputContainer>
            <SC.SearchInput placeholder={'Ссылка на статью...'} />
            <Button backgroundColor={backgroundColorButton}
                    color={colorButton}
                    content={contentButton}
                    handleClick={() => {
                        if(freeRequestsCount === 0 && !isLoggedIn) {
                            dispatch(setAnswerContent('Нет ответа'));
                            return;
                        }
                        console.log('запрос пересказа');
                        dispatch(setAnswerContent('типо ответ'));
                        if(!isLoggedIn) {
                            dispatch(decrementFreeRequestsCount());
                        }
                    }}
            />
        </SC.SearchInputContainer>
    );
};