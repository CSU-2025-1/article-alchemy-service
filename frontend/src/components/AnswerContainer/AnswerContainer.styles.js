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
`;

export const LoginLink = styled(Link)`
    color: #FFFFFF;
    text-decoration: underline;
`;