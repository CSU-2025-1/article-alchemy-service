import styled from "styled-components";
import {Link} from "react-router-dom";


export const Container = styled.div`
    width: 230rem;
    background-color: #FFFFFF;
    border-radius: 50px;
    font-size: 5rem;
    
    padding: 7.5rem 10rem;
    margin-top: 10rem;
    
    height: max-content;
    
    flex-direction: column;
    gap: 5rem;
`;

export const LoginLink = styled(Link)`
    color: #FFFFFF;
    text-decoration: underline;
`;

export const AnswerHeader = styled.h2`
    font-size: 8rem;
    font-weight: 700;
    line-height: 7.5rem;
    letter-spacing: -0.25rem;
`;

export const AnswerHeaderContainer = styled.div`
    display: flex;
    justify-content: space-between;
    align-items: center;
`;

export const TitleHeader = styled.h3`
    font-size: 5rem;
    font-weight: 700;
    line-height: 7.5rem;
    letter-spacing: -0.25rem;
`;

export const TitleContent = styled.p`
    font-size: 5rem;
    font-weight: 400;
    line-height: 7.5rem;
    letter-spacing: -0.25rem;
`;

export const StatusIcon = styled.div`
    width: 5.5rem;
    height: 5rem;
    border-radius: 50%;
`;