import styled from "styled-components";
import {Link} from "react-router-dom";

export const BackLink = styled(Link)`
    text-decoration: none;
    max-width: max-content;
    display: flex;
`;

export const BackLinkText = styled.h2`
    color: white;
    font-size: 20px;
    line-height: 32px;
    max-width: 280rem;
    margin: 0 auto;
`;

export const BackLinkContainer = styled.div`
    max-width: 280rem;
    margin: 0 auto;
`;

export const BackNavComponentContainer = styled.div` 
    width: 100%;
    padding: 4rem 0 4rem 0;
    &:last-child {
        padding-bottom: 20rem;
    }
`;