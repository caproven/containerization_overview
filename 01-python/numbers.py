from time import sleep


def main():
    n = 0
    while True:
        print(f'value: {n}')
        n += 1
        sleep(2)


if __name__ == '__main__':
    main()
