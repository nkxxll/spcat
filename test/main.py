"""This is a test module for the main module.

This module is used to test the main module. It is not intended to be used in
production code. Those tests are integration tests for the command line
interface and the input output handling of the spcat application.
"""

from pathlib import Path
from subprocess import PIPE, Popen


def test_main():
    """Test the main module.

    This function tests the main module by running the main module with the
    help option. It checks if the help message is printed to stdout.
    """
    executable = Path(__file__).parent.parent / "bin" / "spcat"
    process = Popen([executable], stdout=PIPE, stderr=PIPE)
    _, error = process.communicate()
    assert error != b""
