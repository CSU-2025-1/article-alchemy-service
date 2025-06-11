import isPropValid from '@emotion/is-prop-valid';
import { StyleSheetManager as StyleSheetManagerLibrary } from 'styled-components';

const shouldForwardProp = (propName, target) => {
    if (typeof target === 'string') {
        return isPropValid(propName);
    }

    return true;
};

export const StyleSheetManager = ({ children }) => {
    return (
        <StyleSheetManagerLibrary shouldForwardProp={shouldForwardProp}>
            {children}
        </StyleSheetManagerLibrary>
    );
};
