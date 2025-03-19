import * as SC from './Button.styles'

export const Button = ({ backgroundColor, color, content }) => {
    return (
        <SC.ButtonContainer backgroundColor={backgroundColor} color={color}>
            {content}
        </SC.ButtonContainer>
    );
};