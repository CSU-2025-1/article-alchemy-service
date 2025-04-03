import {
    AnswerHeader,
    AnswerHeaderContainer,
    Container,
    StatusIcon, TitleContent, TitleHeader
} from "@/components/AnswerContainer/AnswerContainer.styles.js";
import {useSelector} from "react-redux";

export const AnswerContainer = () => {
    const answerContent = useSelector(state => state.answerContent);
    const freeRequestsCount = useSelector(state => state.freeRequestsCount);
    const isLoggedIn = useSelector(state => state.isLoggedIn);

    let show = (freeRequestsCount !== 0 || isLoggedIn) && (answerContent.status !== 'none');

    let content = <></>;
    if(answerContent && answerContent.status) {
        switch (answerContent.status) {
            case 'pending':
                content =
                    <AnswerHeaderContainer>
                        <AnswerHeader>Ожидание ответа</AnswerHeader>
                        <StatusIcon style={{backgroundColor: '#9175DB'}}></StatusIcon>
                    </AnswerHeaderContainer>;
                break;
            case 'failed':
                content =
                    <AnswerHeaderContainer>
                        <AnswerHeader>Что-то пошло не так :(</AnswerHeader>
                        <StatusIcon style={{backgroundColor: '#FF5959'}}></StatusIcon>
                    </AnswerHeaderContainer>;
                break;
            case 'completed':
                content =
                    <>
                        <AnswerHeaderContainer>
                            <AnswerHeader>{answerContent.body.main_title}</AnswerHeader>
                            <StatusIcon style={{backgroundColor: '#1CED00'}}></StatusIcon>
                        </AnswerHeaderContainer>
                        {answerContent.body.data.map((item, index) => (
                            <div key={index}>
                                <TitleHeader key={`title-${index}`}>{item.title}</TitleHeader>
                                {item.points.map((point, pointIndex) => (
                                    <TitleContent key={pointIndex}>
                                          •  {point}
                                    </TitleContent>
                                ))}
                            </div>
                        ))}
                    </>;
                break;
        }
    }

    return (
        <Container style={{display: show ? 'flex' : 'none'}}>
            {content}
        </Container>
    );
};