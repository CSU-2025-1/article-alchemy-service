import { SearchInput } from '@/components/SearchInput';
import {NavBar} from "./components/NavBar";
import * as SC from './SummaryPage.styles';
import {Subtitle, Title} from "./SummaryPage.styles";
import {AnswerContainer} from "@/components/AnswerContainer/AnswerContainer.jsx";
import {RequestsCount} from "@/components/RequestsCountContainer/RequestsCount.jsx";
import {NavBarLinks} from "@/pages/SummaryPage/components/NavBarLinks/NavBarLinks.jsx";
import {useDispatch} from "react-redux";
import {useNavigate} from "react-router-dom";
import {useEffect} from "react";
import {getUserInfoRequest, refreshAccessTokenRequest} from "@/api/api.js";
import {setIsLoggedIn, setIsLogin, setProfileEmail, setProfileName} from "@/store/appSlice.js";
import {ROUTES} from "@/app/Router/routes.js";

export const SummaryPage = () => {
    const dispatch = useDispatch();
    const navigate = useNavigate();

    useEffect(() => {
        (async () => {
            const refreshToken = localStorage.getItem('refreshToken');
            if (refreshToken) {
                if (await refreshAccessTokenRequest(refreshToken)) {
                    const userInfo = await getUserInfoRequest();
                    dispatch(setIsLoggedIn(true));
                    dispatch(setProfileName(userInfo.username));
                    dispatch(setProfileEmail(userInfo.email));
                }
                else {
                    alert('Срок сессии истёк. Войдите заново.');
                    dispatch(setIsLogin(true));
                    navigate(ROUTES.auth);
                }
            }
        })();
    }, []);

    return (
        <SC.Wrapper>
            <NavBar>
                <NavBarLinks />
            </NavBar>
            <SC.SummaryContainer>
                <Title>
                    Создать краткий пересказ статьи прямо сейчас
                </Title>
                <Subtitle>
                    Введите ссылку на статью и получите краткий пересказ
                </Subtitle>

                <SearchInput />
                <AnswerContainer />
                <RequestsCount />

            </SC.SummaryContainer>
        </SC.Wrapper>

    );
};