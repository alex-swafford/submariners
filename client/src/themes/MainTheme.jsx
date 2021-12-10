import { createTheme } from "@material-ui/core"

const MainTheme = createTheme({
    palette: {
        primary: {
            main: '#00709f'
        },
        secondary: {
            main: '#00f070'
        }
    }
})

MainTheme.spacing(2)

export default MainTheme