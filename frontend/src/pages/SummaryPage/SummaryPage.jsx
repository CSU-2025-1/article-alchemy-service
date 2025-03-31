import { SearchInput } from '@/components/SearchInput';
import {NavBar} from "./components/NavBar";
import * as SC from './SummaryPage.styles';
import {Subtitle, Title} from "./SummaryPage.styles";
import {AnswerContainer} from "@/components/AnswerContainer/AnswerContainer.jsx";
import {RequestsCount} from "@/components/RequestsCountContainer/RequestsCount.jsx";
import {NavBarLinks} from "@/pages/SummaryPage/components/NavBarLinks/NavBarLinks.jsx";

export const SummaryPage = () => {
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