module Main exposing (Model, Msg, init, subscriptions, update, view)

import Api.Object
import Api.Query
import Api.Scalar
import Api.ScalarCodecs
import Browser
import Graphql.Http
import Graphql.Operation exposing (RootQuery)
import Html exposing (..)
import RemoteData


main : Program () Model Msg
main =
    Browser.document
        { init = init
        , view = view
        , update = update
        , subscriptions = subscriptions
        }


type alias Model =
    { id : Int
    }


makeRequest : Cmd Msg
makeRequest =
    Api.Query.user (Api.Query.UserRequiredArguments (Api.Scalar.Id "1")) Api.Object.User
        |> Graphql.Http.queryRequest "http://localhost:8080/graphql"
        |> Graphql.Http.send (RemoteData.fromResult >> GotResponse)


init : () -> ( Model, Cmd Msg )
init _ =
    ( Model 0, Cmd.none )


type Msg
    = GotResponse Model


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        GotResponse response ->
            ( response, Cmd.none )


subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none


view : Model -> Browser.Document Msg
view model =
    { title = "MeetMeUp App"
    , body =
        [ div []
            [ text "New Document" ]
        ]
    }
