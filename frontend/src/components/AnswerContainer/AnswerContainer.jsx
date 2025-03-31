import {Container} from "@/components/AnswerContainer/AnswerContainer.styles.js";
import {useSelector} from "react-redux";


export const AnswerContainer = () => {
    const answerContent = useSelector(state => state.answerContent);
    const freeRequestsCount = useSelector(state => state.freeRequestsCount);
    const isLoggedIn = useSelector(state => state.isLoggedIn);

    let show = (freeRequestsCount !== 0 || isLoggedIn) && answerContent;

    return (
        <Container style={{display: show ? 'block' : 'none'}}>
            {answerContent}
        </Container>
    );
};