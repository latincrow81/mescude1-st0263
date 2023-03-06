from dynaconf import Dynaconf


class Config(Dynaconf):
    def force_for_testing(self) -> None:
        self.configure(FORCE_ENV_FOR_DYNACONF="testing")


main_config = Config(environments=True,
                     settings_files=["settings.toml"],
                     load_dotenv=False,
                     env_switcher='ENV')


def get_config() -> Config:
    return main_config
