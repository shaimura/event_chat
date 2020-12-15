import axios from 'axios';

// axios内でrouterがうまく機能しないため、コメントアウト
// import router from '../../router';

// 画面をリロードしても"state"にデータを残すためのプラグイン
import createPersistedState from "vuex-persistedstate";



const state = {
    idToken: null
};

const getters = {
    idToken: state => state.idToken
};

const mutations = {
    updateIdToken(state, idToken) {
        state.idToken = idToken;
    }
};


const actions = {
    async autoLogin({ commit, dispatch }) {
        const idToken = localStorage.getItem('idToken');
        if (!idToken) return;
        const now = new Date();
        // トークンの有効期限を取得する
        const exporyTimeMs = localStorage.getItem('exporyTimeMs');
        // 今の時間と有効期限の時間を比較して"ture"か"folse"を取得する
        const isExprired = now.getTime() >= exporyTimeMs;
        const userid = localStorage.getItem('userid');
        if (isExprired) {
            // 時間経過でログアウトする場合は、時間判定のif文を追加し、ローカルストレージをクリアする

            await dispatch('refreshIdToken', userid);
        } else {
            const expiresInMs = exporyTimeMs - now.getTime();
            setTimeout(() => {
                dispatch('refreshIdToken', userid)
            }, expiresInMs)
            commit('updateIdToken', idToken);
        }
    },
    signin({ dispatch }, params) {
        axios.post('/signin', params).then(response => {
            if (!response.data.message) {
                alert("ログインに失敗しました")
            } else {
                const user = response.data.user;
                const userparams = new URLSearchParams();
                userparams.append('userid', user.ID);
                console.log(user);
                console.log(userparams);
                axios.post('/refreshidtoken', userparams).then(returntoken => {
                    dispatch('setAuthDate', {
                        idToken: returntoken.data.Accesstoken,
                        expiresIn: returntoken.data.Expirationdate,
                        userName: returntoken.data.Username,
                        userId: returntoken.data.UserID
                    });
                    // router.push('/');
                    location.href = "/";
                    alert(response.data.message);
                })
            }
        }).catch(error => {
            console.log(error.response);
        });
    },
    logout({ commit }) {
        commit('token/updateIdToken', null);
        localStorage.clear();
        // router.replace('/login');
        location.href = "/login";

    },
    async refreshIdToken({ dispatch }, userid) {
        const params = new URLSearchParams();
        params.append('userid', userid);
        await axios.post('/refreshidtoken', params).then(returntoken => {
            dispatch('token/setAuthDate', {
                idToken: returntoken.data.Accesstoken,
                expiresIn: returntoken.data.Expirationdate,
                userName: returntoken.data.Username,
                userId: returntoken.data.UserID
            })
        }).catch(error => {
            console.log(error);
        })
    },
    signup({ dispatch }, params) {
        axios.post('/signup', params).then(response => {
            if (!response.data.errormessage) {
                const user = response.data.user;
                const userparams = new URLSearchParams();
                userparams.append('id', user.ID);
                dispatch('setAuthDate', {
                    idToken: response.data.accesstoken.Accesstoken,
                    userName: response.data.accesstoken.Username,
                    expiresIn: response.data.accesstoken.Expirationdata,
                    userId: response.data.accesstoken.UserID
                });
                
                alert("ログインしました");
                // router.push('/');
                location.href = "/";
                } else {
                let errorMessage = response.data.errormessage;
                errorMessage = errorMessage.join("");
                console.log(errorMessage);
                alert(errorMessage);
            }
        }).catch(error => {
            console.log(error);
        });        
    },
    gettoken(){},
    setAuthDate({ commit, dispatch }, authData){
        // １時間m秒
        const settime = 3600000;
        const exporyTimeMs = authData.expiresIn;

        // ログインした時にトークンを与える
        commit('updateIdToken', authData.idToken);

        // ローカルストレージにデータを登録する
        localStorage.setItem('idToken', authData.idToken);
        localStorage.setItem('userName', authData.userName);
        localStorage.setItem('exporyTimeMs', authData.expiresIn)
        localStorage.setItem('userID', authData.userId);
        
        setTimeout(() => {
            dispatch('refreshIdToken', authData.userId)
        }, exporyTimeMs + settime);
    }
};


export default {
    namespaced: true,
    state: state,
    getters: getters,
    mutations: mutations,
    actions: actions,
    plugins: [createPersistedState()]
};
