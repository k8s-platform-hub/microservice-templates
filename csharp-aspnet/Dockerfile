FROM microsoft/dotnet:2.0.4-sdk-2.1.3-jessie
RUN mkdir /src
WORKDIR /src
ADD app/src/ /src
RUN dotnet restore
RUN dotnet build
CMD dotnet run
