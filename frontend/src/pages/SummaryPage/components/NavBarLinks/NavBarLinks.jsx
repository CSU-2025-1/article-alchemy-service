import {Link, useNavigate} from "react-router-dom";
import {setAnswerContent, setIsLoggedIn, setIsLogin, setProfileName} from "@/store/appSlice.js";
import {Button} from "@/components/Button/index.js";
import {useDispatch, useSelector} from "react-redux";
import {ROUTES} from "@/app/Router/routes.js";
import {logoutRequest} from "@/api/api.js";

export const NavBarLinks = () => {
    const dispatch = useDispatch();
    const navigate = useNavigate();
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
                            handleClick={async () => {
                                const response = await logoutRequest();
                                console.log('ответ с серва при разлогине', response);
                                if(response){
                                    localStorage.removeItem('refreshToken');
                                    localStorage.removeItem('accessToken');
                                    dispatch(setIsLoggedIn(false));
                                    dispatch(setAnswerContent({data: null, status: 'none'}));
                                    dispatch(setIsLogin(true));
                                    dispatch(setProfileName(''));
                                }
                                else {
                                    await navigate(ROUTES.root);
                                    alert('Не удалось выйти...');
                                }
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