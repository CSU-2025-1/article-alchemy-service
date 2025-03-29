import { Button } from '@/components/Button';
import { SearchInput } from '@/components/SearchInput';
import {NavBar} from "./components/NavBar";
import * as SC from './SummaryPage.styles'
import {AuthForm} from "@/components/AuthForm/index.js";
export const SummaryPage = () => {
    return (
        <SC.Wrapper>
            <NavBar>
                <Button backgroundColor={'white'} color={'var(--color-grape)'} content={'Вход'}></Button>
                <Button backgroundColor={'var(--color-grape)'} color={'white'} content={'Регистрация'}></Button>
            </NavBar>
            <SC.SummaryContainer>
                <SearchInput></SearchInput>
                <AuthForm></AuthForm>
            </SC.SummaryContainer>
        </SC.Wrapper>

    );
};