FROM ubuntu:latest

# install python
RUN apt update && \
    # without env, tzdata dep prompts for timezone
    DEBIAN_FRONTEND=noninteractive apt install -y python3

RUN useradd \
    --system \
    --uid 1001 \
    --create-home \
    myuser

USER myuser

WORKDIR /app

COPY ./numbers.py .

CMD ["python3", "numbers.py"]
