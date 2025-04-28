import pytest

if __name__ == "__main__":
    pytest.main([
        "-v",
        "--html=report.html",
        "--capture=no",
        "autotests/"
    ])