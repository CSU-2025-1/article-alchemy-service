import {Button} from "@/components/Button/index.js";
import {ROUTES} from "@/app/Router/routes.js";
import {setAnswerContent} from "@/store/appSlice.js";
import {AnswerContainer} from "@/components/AnswerContainer/AnswerContainer.jsx";
import React, {useEffect, useState} from "react";
import {useDispatch} from "react-redux";
import {BackNavComponent} from "@/components/BackNavigation/BackNavComponent.jsx";
import {Title} from "@/pages/SummaryPage/SummaryPage.styles.js";
import {
    HistoryContainer,
    HistoryElementContainer, HistoryTitle
} from "@/pages/HistoryPage/components/HistoryPageContent/HistoryPageContent.styles.js";
import {getHistory} from "@/api/api.js";
import {StatusIcon} from "@/components/AnswerContainer/AnswerContainer.styles.js";

export const HistoryPageContent  = () => {
    const dispatch = useDispatch();
    const [history, setHistory] = useState([]);
    const [isHistoryShow, setIsHistoryShow] = useState(true);
    const [historyLimit, setHistoryLimit] = useState(5);

    useEffect(() => {
        (async () => {
            const result = await getHistory();
            await setHistory(result.items);
        })();
    }, []);

    const button = historyLimit > history.length
        ? <></>
        : <Button content={'Еще'}
                  handleClick={() => {
                      setHistoryLimit(historyLimit + 5);
                  }}
                  backgroundColor={'#FFFFFF'}
                  color={'black'}/>;

    if (isHistoryShow) {
        return (
            <>
                <BackNavComponent route={ROUTES.root}
                                  onClick={() => {
                                      dispatch(setAnswerContent({body: null, status: 'none'}));
                                  }}/>
                <Title>История запросов</Title>
                <HistoryContainer style={{display: isHistoryShow ? "flex" : "none"}}>
                    {history.slice(0, historyLimit).map((item, index) => (
                        <HistoryElementContainer key={index}>
                            <HistoryTitle>
                                {item.status === 'completed'
                                ?
                                    <>
                                        <StatusIcon style={{backgroundColor: '#1CED00'}}></StatusIcon>
                                        {JSON.parse(item.data).main_title}
                                    </>
                                :item.status === 'pending'
                                ?
                                    <>
                                        <StatusIcon style={{backgroundColor: '#9175DB'}}></StatusIcon>
                                        Ожидание ответа
                                    </>
                                :
                                    <>
                                        <StatusIcon style={{backgroundColor: '#FF5959'}}></StatusIcon>
                                        Что то пошло не так :(
                                    </>
                                }
                            </HistoryTitle>
                            <Button content={'Просмотреть'}
                                    handleClick={() => {
                                        setIsHistoryShow(false);
                                        dispatch(setAnswerContent({
                                            body: JSON.parse(item.data),
                                            status: item.status
                                        }));
                                    }}
                                    backgroundColor={'#4870F1'}/>
                        </HistoryElementContainer>
                    ))}
                    {button}
                </HistoryContainer>
            </>
        );
    }

    return (
        <>
            <BackNavComponent route={ROUTES.root}
                              onClick={() => {
                                  dispatch(setAnswerContent({body: null, status: 'none'}));
                              }}
                              content={'На главную'}/>
            <BackNavComponent route={ROUTES.history}
                              onClick={() => {
                                  setIsHistoryShow(true);
                              }}
                              content={'К истории'}/>
            <Title>История запросов</Title>
            <AnswerContainer />
        </>
    );
};