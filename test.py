import os


def run():
    # exec "python host.py" and capture stdout
    import subprocess
    p = subprocess.Popen(["python", "host.py"], stdout=subprocess.PIPE)
    stdout, stderr = p.communicate()

    print(stdout)

    # check stdout contains "a1: 1"
    assert b"a1: 1" in stdout
    assert b"a2: 2" in stdout
    assert b"42" in stdout

if __name__ == '__main__':
    run()