import {NavBar} from "@/pages/SummaryPage/components/NavBar/index.js";
import React from "react";
import {Wrapper} from "@/pages/SummaryPage/SummaryPage.styles.js";
import {HistoryPageContent} from "@/pages/HistoryPage/components/HistoryPageContent/HistoryPageContent.jsx";
import {NavBarLinks} from "@/pages/SummaryPage/components/NavBarLinks/NavBarLinks.jsx";

export const HistoryPage = () => {

    return (
        <Wrapper>
            <NavBar>
                <NavBarLinks />
            </NavBar>
            <HistoryPageContent />
        </Wrapper>
    );
};