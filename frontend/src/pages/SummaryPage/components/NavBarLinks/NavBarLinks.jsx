import {Link} from "react-router-dom";
import {setAnswerContent, setIsLoggedIn, setIsLogin, setProfileName} from "@/store/appSlice.js";
import {Button} from "@/components/Button/index.js";
import {useDispatch, useSelector} from "react-redux";
import {ROUTES} from "@/app/Router/routes.js";

export const NavBarLinks = () => {
    const dispatch = useDispatch();
    const isLoggedIn = useSelector(state => state.isLoggedIn);

    if(isLoggedIn){
        return (
            <>
                <Link to={ROUTES.history}>
                    <Button backgroundColor={'#151718'}
                            color={'white'}
                            content={'История'}>
                    </Button>
                </Link>
                <Link to='/auth'>
                    <Button backgroundColor={'white'}
                            color={'var(--color-grape)'}
                            content={'Выйти'}
                            handleClick={() => {
                                // ...
                                dispatch(setIsLoggedIn(false));
                                dispatch(setAnswerContent(null));
                                dispatch(setIsLogin(true));
                                dispatch(setProfileName(''));
                            }}>
                    </Button>
                </Link>
            </>
        );
    }

    return (
        <>
            <Link to='/auth'
                  onClick={() => {
                      dispatch(setIsLogin(true));
                  }}>
                <Button backgroundColor={'white'}
                        color={'var(--color-grape)'}
                        content={'Вход'}>
                </Button>
            </Link>
            <Link to='/auth'
                  onClick={() => {
                      dispatch(setIsLogin(false));
                  }}>
                <Button backgroundColor={'var(--color-grape)'}
                        color={'white'}
                        content={'Регистрация'}>
                </Button>
            </Link>
        </>
    );
};