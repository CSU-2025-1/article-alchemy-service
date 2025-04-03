import {NavBar} from "../SummaryPage/components/NavBar";
import * as SC from '../SummaryPage/SummaryPage.styles';
import {AuthForm} from "@/components/AuthForm/index.js";
import {ROUTES} from "@/app/Router/routes.js";
import {BackNavComponent} from "@/components/BackNavigation/BackNavComponent.jsx";

export const AuthPage = () => {
    return (
        <SC.Wrapper>
            <NavBar />
            <BackNavComponent route={ROUTES.root}/>
            <AuthForm />
        </SC.Wrapper>
    );
};