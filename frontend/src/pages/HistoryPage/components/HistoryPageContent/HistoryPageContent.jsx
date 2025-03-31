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

export const HistoryPageContent  = () => {
    const dispatch = useDispatch();
    const [history, setHistory] = useState([]);
    const [isHistoryShow, setIsHistoryShow] = useState(true);
    const [historyLimit, setHistoryLimit] = useState(5);

    useEffect(() => {
        (async () => {
            await setHistory([
                {header: '1dsgsgsgfgdggdfsgfsddsgfsgfd'},
                {header: '2sdfggfgfdgdgfsdgfdgfgdffdg'},
                {header: '3dfsgsfgdfsgsdgfdfgdgffgdsfgd'},
                {header: '4dsfdfgdfgfsdsfdgfgdsgsdfs'},
                {header: '5dsfdfgdfgfsdsfdgfgdsgsdfs'},
                {header: '6dsfdfgdfgfsdsfdgfgdsgsdfs'},
            ]);
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
                                      dispatch(setAnswerContent(null));
                                  }}/>
                <Title>История запросов</Title>
                <HistoryContainer style={{display: isHistoryShow ? "flex" : "none"}}>
                    {history.slice(0, historyLimit).map((item, index) => (
                        <HistoryElementContainer key={index}>
                            <HistoryTitle>{item.header}</HistoryTitle>
                            <Button content={'Просмотреть'}
                                    handleClick={() => {
                                        setIsHistoryShow(false);
                                        dispatch(setAnswerContent(item.header));
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
                                  dispatch(setAnswerContent(null));
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