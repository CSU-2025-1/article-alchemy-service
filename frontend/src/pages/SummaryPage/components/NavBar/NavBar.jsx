import * as SC from './NavBar.styles'
import {MainIcon} from "@/components/Icons/index.js";

export const NavBar = ( { children } ) => {
    return (
        <SC.NavContainer>
            <SC.NavBar>
                <SC.NavElem>
                    <MainIcon/>
                </SC.NavElem>
                <SC.NavElem>
                    {children}
                </SC.NavElem>
            </SC.NavBar>
        </SC.NavContainer>
    );
};