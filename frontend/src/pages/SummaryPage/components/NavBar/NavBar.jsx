import * as SC from './NavBar.styles';
import {MainIcon} from "@/components/Icons/index.js";
import {ProfileText} from "./NavBar.styles";
import {useSelector} from "react-redux";

export const NavBar = ( { children } ) => {
    const profileName = useSelector(state => state.profileName);
    return (
        <SC.NavContainer>
            <SC.NavBar>
                <SC.NavElem>
                    <MainIcon/>
                    <ProfileText>
                        {profileName}
                    </ProfileText>
                </SC.NavElem>
                <SC.NavElem>
                    {children}
                </SC.NavElem>
            </SC.NavBar>
        </SC.NavContainer>
    );
};