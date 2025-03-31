import {RequestsCountContainer} from "@/components/RequestsCountContainer/RequestsCount.styles.js";
import {useDispatch, useSelector} from "react-redux";
import {LoginLink} from "@/components/AnswerContainer/AnswerContainer.styles.js";
import {setIsLogin} from "@/store/appSlice.js";

export const RequestsCount = () => {
    const dispatch = useDispatch();
    const answerContent = useSelector(state => state.answerContent);
    const freeRequestsCount = useSelector(state => state.freeRequestsCount);
    const isLoggedIn = useSelector(state => state.isLoggedIn);

    let firstWord = 'Осталось';
    let requestWord;
    switch (freeRequestsCount) {
        case 0:
            requestWord = 'запросов';
            break;
        case 1:
            requestWord = 'запрос';
            firstWord = 'Остался';
            break;
        case 5:
            requestWord = 'запросов';
            break;
        default:
            requestWord = 'запроса';
            break;
    }

    return (
        <RequestsCountContainer style={{display: isLoggedIn
                ? 'none'
                : answerContent ? 'block' : 'none'}}>
            {firstWord} <strong>{freeRequestsCount} бесплатных</strong> {requestWord}. Для увеличения количества запросов <LoginLink to={'auth'} onClick={() => dispatch(setIsLogin(true))}>войдите в систему.</LoginLink>
        </RequestsCountContainer>
    );
};