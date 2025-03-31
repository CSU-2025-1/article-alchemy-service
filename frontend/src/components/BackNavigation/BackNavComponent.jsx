import {
    BackLink,
    BackLinkContainer,
    BackLinkText,
    BackNavComponentContainer
} from "@/components/BackNavigation/BackNavComponent.styles.js";


export const BackNavComponent = ({route, onClick, content}) => {
    return (
        <BackNavComponentContainer onClick={onClick}>
            <BackLinkContainer>
                <BackLink to={route}>
                    <BackLinkText>{content ? content : '← Назад'}</BackLinkText>
                </BackLink>
            </BackLinkContainer>
        </BackNavComponentContainer>
    );
};